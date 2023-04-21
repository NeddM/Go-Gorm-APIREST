package routes

import (
	"encoding/json"
	"net/http"

	"github.com/NeddM/go-gorm-apirest/db"
	"github.com/NeddM/go-gorm-apirest/models"
	"github.com/gorilla/mux"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task []models.Task
	db.DB.Find(&task)

	if len(task) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data"))
		return
	}

	json.NewEncoder(w).Encode(task)
	w.Write([]byte("Get task"))
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func PostTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
	w.Write([]byte("Task created"))
}

func DeleteTasksHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	db.DB.Delete(&task)
	w.Write([]byte("Task deleted"))
}

func EditTasksHandler(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	var task models.Task
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&newTask)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Where("id = ?", task.ID).Updates(&newTask)
	w.Write([]byte("Task updated"))
}
