package models

import (
	"time"
)

type Task struct {
	ID           int       `json:"id"`
	TaskName     string    `json:"task_name"`
	ScheduleTime time.Time `json:"schedule_time"`
	Command      string    `json:"command"`
	Status       string    `json:"status"`
	Priority     string    `json:"priority"`
}
