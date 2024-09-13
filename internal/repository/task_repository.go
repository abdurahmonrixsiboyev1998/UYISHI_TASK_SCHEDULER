package repository

import (
	"database/sql"
	"log"
	"scheduler/internal/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := `INSERT INTO tasks (task_name, schedule_time, command, status, priority) 
              VALUES ($1, $2, $3, 'pending', $4) RETURNING id`
	err := r.DB.QueryRow(query, task.TaskName, task.ScheduleTime, task.Command, task.Priority).Scan(&task.ID)
	if err != nil {
		log.Println("Failed to create task:", err)
		return err
	}
	return nil
}

func (r *TaskRepository) GetPendingTasks() ([]models.Task, error) {
	query := `SELECT id, task_name, schedule_time, command, status, priority FROM tasks WHERE status = 'pending'`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.TaskName, &task.ScheduleTime, &task.Command, &task.Status, &task.Priority); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) UpdateTaskStatus(id int, status string) error {
	query := `UPDATE tasks SET status = $1 WHERE id = $2`
	_, err := r.DB.Exec(query, status, id)
	if err != nil {
		return err
	}
	return nil
}
