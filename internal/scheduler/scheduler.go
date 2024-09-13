// scheduler/scheduler.go
package scheduler

import (
	"log"
	"time"

	"scheduler/internal/models"
	"scheduler/internal/repository" 
)

type Scheduler struct {
	Repo *repository.TaskRepository
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-ticker.C:
			s.RunPendingTasks()
		}
	}
}

func (s *Scheduler) RunPendingTasks() {
	tasks, err := s.Repo.GetPendingTasks()
	if err != nil {
		log.Println("Failed to retrieve pending tasks:", err)
		return
	}

	for _, task := range tasks {
		if time.Now().After(task.ScheduleTime) {
			s.ExecuteTask(task)
		}
	}
}

func (s *Scheduler) ExecuteTask(task models.Task) {
	log.Println("Executing task:", task.TaskName)
	err := s.Repo.UpdateTaskStatus(task.ID, "completed")
	if err != nil {
		log.Println("Failed to update task status:", err)
	}
}
