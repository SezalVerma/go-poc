package routes

import (
	"github.com/gorilla/mux"
	"github.com/sezalverma/go-poc/controllers"
)

//import goes here

func UserRoute(router *mux.Router) {
    //other routes goes here
    // router.GET("/user/:userId", controllers.GetAUser()) //add this

	router.HandleFunc("/api/users" , controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}" , controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/users/{id}",controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/aadhar/{aadhar}", controllers.GetUserByAadhar).Methods("GET")
	router.HandleFunc("/api/users/phone/{phone}",controllers.GetUserByPhone).Methods("GET")
	router.HandleFunc("/api/users/pan/{pan}",controllers.GetUserByPan).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
}