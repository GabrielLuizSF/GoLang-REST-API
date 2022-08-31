package server


import (
	
	"github.com/gorilla/mux"
	"user_api-golang/controllers"
)

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/api/users", controllers.GetAllUser).Methods("GET")
	router.HandleFunc("/api/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user/edit/{id}", controllers.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/api/user/delete/{id}", controllers.DeleteUserByID).Methods("DELETE")
}
