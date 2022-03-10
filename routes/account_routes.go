package routes

import (
	"github.com/gorilla/mux"
	"github.com/sezalverma/go-poc/controllers"
)

func AccountRoute(router *mux.Router) {
    //other routes goes here
    // router.GET("/user/:userId", controllers.GetAUser()) //add this

	router.HandleFunc("/api/accounts" , controllers.GetAllAccounts).Methods("GET")
	// router.HandleFunc("/api/users/{id}" , controllers.GetUserById).Methods("GET")
	// router.HandleFunc("/api/users/aadhar/{aadhar}", controllers.GetUserByAadhar).Methods("GET")
	// router.HandleFunc("/api/users/phone/{phone}",controllers.GetUserByPhone).Methods("GET")
	// router.HandleFunc("/api/users/pan/{pan}",controllers.GetUserByPan).Methods("GET")
	router.HandleFunc("/api/accounts", controllers.CreateAccount).Methods("POST")
}