package router

import (
	"rest-crud-api-go-lang/controller"

	"github.com/gorilla/mux"
)

func InitaliseHandlers(router *mux.Router) {
	router.HandleFunc("/users",
		controller.CreateUser).Methods("POST")
	router.HandleFunc("/users",
		controller.GetAllUser).Methods("GET")
	router.HandleFunc("/users/{id}",
		controller.GetUserByID).Methods("GET")
	router.HandleFunc("/users",
		controller.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/users/{id}",
		controller.DeletUserByID).Methods("DELETE")
}
