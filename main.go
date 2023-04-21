package main

import (
	"net/http"

	"github.com/NeddM/go-gorm-apirest/db"
	"github.com/NeddM/go-gorm-apirest/models"
	"github.com/NeddM/go-gorm-apirest/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	// Users routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", routes.EditUsersHandler).Methods("PUT")

	// Task routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTasksHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTasksHandler).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", routes.EditTasksHandler).Methods("PUT")

	http.ListenAndServe(":3000", r)
}
