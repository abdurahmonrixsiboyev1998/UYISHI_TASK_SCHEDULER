package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"scheduler/internal/models"
	"scheduler/internal/repository"
)

type TaskHandler struct {
	Repo *repository.TaskRepository
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	task.ScheduleTime, err = time.Parse(time.RFC3339, r.FormValue("schedule_time"))
	if err != nil {
		http.Error(w, "Invalid schedule time", http.StatusBadRequest)
		return
	}

	err = h.Repo.CreateTask(&task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}
