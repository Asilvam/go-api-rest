package controllers

import (
	"encoding/json"
	"go-api-rest/src/models"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool                                 `json:"success"`
	Data    []models.Employees_emailnotification `json:"data"`
	Errors  []string                             `json:"errors"`
}

func GetTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var data Data
	var todo models.Employees_emailnotification
	var success bool
	todo, success = models.Get(id)
	if !success {
		data.Success = false
		data.Errors = append(data.Errors, "employees_emailnotification not found")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetTodos(w http.ResponseWriter, req *http.Request) {
	var todos []models.Employees_emailnotification = models.GetAll()

	var data = Data{true, todos, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
