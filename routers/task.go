package routers

import (
	"github.com/gorilla/mux"
	"github.com/hecomp/taskmanager/controllers"
)

func SetTaskRouters(router *mux.Router) *mux.Router  {
	router.HandleFunc("/tasks", controllers.Tasks.Create).Methods("POST")
	router.HandleFunc("/tasks", controllers.Tasks.Create).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Create).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Create).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.Tasks.Create).Methods("DELETE")
	return router
}
