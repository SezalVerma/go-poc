package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sezalverma/go-poc/configs"
	"github.com/sezalverma/go-poc/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connect to mongoDB with db.setup
// var db = configs.ConnectDB()
var user_collection = configs.ConnectDB().Collection("users")

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var users []models.User

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := user_collection.Find(context.TODO(), bson.M{})

	if err != nil {
		configs.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var user models.User
		// & character returns the memory address of the following variable.
		err := cur.Decode(&user) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	if len(users) >0 {
        w.Write([]byte( strconv.Itoa(len(users)) + " Users Found \n" )) 
		json.NewEncoder(w).Encode(users) // encode similar to serialize process.
	} else {
	   w.Write([]byte("No User Found \n"))
	}

	
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := user_collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		// configs.GetError(err, w)
		w.Write([]byte("No user found with this Id \n"))
		// return
	}else{
        w.Write([]byte("User Found \n"))

	    json.NewEncoder(w).Encode(user)
	}
     
	
}

func GetUserByAadhar(w http.ResponseWriter , r *http.Request){
    // set Header
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var params = mux.Vars(r)

	aadhar,_:= strconv.Atoi(params["aadhar"])

	filter := bson.M{"aadhar": aadhar}
	err := user_collection.FindOne(context.TODO(),filter).Decode(&user)

	if err != nil {
		// configs.GetError(err, w)
		w.Write([]byte("No user found with given Aadhar Number \n"))
		// return
	}else{
		w.Write([]byte("User found  \n"))
        json.NewEncoder(w).Encode(user)
	}

	
}

func GetUserByPan(w http.ResponseWriter , r *http.Request){
    // set Header
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var params = mux.Vars(r)

	pan:= params["pan"]

	filter := bson.M{"pan": pan}
	err := user_collection.FindOne(context.TODO(),filter).Decode(&user)

	if err != nil {
		// configs.GetError(err, w)
		w.Write([]byte("No user found with given Pan Number \n"))
		// return
	}else{
		w.Write([]byte("User found  \n"))
		json.NewEncoder(w).Encode(user)
	}

	
}

func GetUserByPhone(w http.ResponseWriter , r *http.Request){
    
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var params= mux.Vars(r)

    fmt.Println(params["phone"])

	phone,_:= strconv.Atoi(params["phone"])

	filter := bson.M{"phone": phone}
	err := user_collection.FindOne(context.TODO(), filter).Decode(&user)
   
	if err!= nil{
		// configs.GetError(err, w)
		// return
		w.Write([]byte("No user registered with this Phone Number" + strconv.Itoa(phone) + "\n "))
	}else{
		w.Write([]byte("User Found \n"))
        json.NewEncoder(w).Encode(user)
	}

	
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entered")
	w.Header().Set("Content-Type", "application/json")

	var user models.User
    
	fmt.Println("before")

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&user)

	// 
    fmt.Println("after")

	// insert our book model.
	result, err := user_collection.InsertOne(context.TODO(), user)

    fmt.Print("here")

	if err != nil {
		fmt.Print("in err")
		configs.GetError(err, w)
		w.Write([]byte("User could not be created"))
		// return
	}else{
		w.Write([]byte("User created successfully !!"))
		
}
fmt.Print("out here")
          json.NewEncoder(w).Encode(result)
	}

	