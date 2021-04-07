package main

import (
	"fmt"
	"go-api-rest/src/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var port string = "8080"
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	apiRouter.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)
}
