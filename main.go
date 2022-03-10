package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sezalverma/go-poc/configs"
	"github.com/sezalverma/go-poc/routes"
)

//Connect to mongoDB with db.setup
// var db =

func main(){
	// init router
	router := mux.NewRouter()

	// handle routes
	routes.UserRoute(router) //add this
    routes.AccountRoute(router) //add this

    configs.ConnectDB()
    

	// set port address
	log.Fatal(http.ListenAndServe(":8000",router))
}
