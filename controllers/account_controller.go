package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sezalverma/go-poc/configs"
	"github.com/sezalverma/go-poc/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


var accounts_collection = configs.DB.Collection("accounts")

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var account models.Account

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&account)

	// insert our book model.
	result, err := accounts_collection.InsertOne(context.TODO(), account)

	if err != nil {
		configs.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var accounts []models.Account

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := accounts_collection.Find(context.TODO(), bson.M{})

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
		var account models.Account
		// & character returns the memory address of the following variable.
		err := cur.Decode(&account) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		accounts = append(accounts, account)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(accounts) // encode similar to serialize process.
}

func GetAllAccountsByAadhar(w http.ResponseWriter , r *http.Request){
    // set Header
	w.Header().Set("Content-Type", "application/json")

	var accounts []models.Account
	var params = mux.Vars(r)

	aadhar,_:= strconv.Atoi(params["aadhar"])

	filter := bson.M{"aadhar": aadhar}

	cur, err := accounts_collection.Find(context.TODO(), filter)

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
		var account models.Account
		// & character returns the memory address of the following variable.
		err := cur.Decode(&account) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		accounts = append(accounts, account)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(accounts) // encode similar to serialize process.
}

func GetAccountById(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var account models.Account
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := accounts_collection.FindOne(context.TODO(), filter).Decode(&account)

	if err != nil {
		configs.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(account)
}
