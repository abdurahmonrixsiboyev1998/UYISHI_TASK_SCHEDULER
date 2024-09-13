package main

import (
	"log"
	"net/http"
	"os"
	"scheduler/config"
	"scheduler/internal/repository"
	"scheduler/internal/scheduler"
	"scheduler/internal/utils"

	handler "scheduler/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	utils.Example()

	repo := &repository.TaskRepository{DB: config.DB}
	taskHandler := &handler.TaskHandler{Repo: repo}

	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")

	sched := scheduler.Scheduler{Repo: repo}
	go sched.Start()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
