package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	dbconfig "rest-crud-api-go-lang/config"
	"rest-crud-api-go-lang/entity"
	"strconv"

	"github.com/gorilla/mux"
)

//Fetch all user get all User data
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var users []entity.User
	dbconfig.Connector.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

//Get user by id returns User with specific ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var user entity.User
	dbconfig.Connector.First(&user, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestBody, &user)

	dbconfig.Connector.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

//Update user by id
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var user entity.User
	json.Unmarshal(requestBody, &user)
	dbconfig.Connector.Save(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//Delet user by id
func DeletUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var user entity.User
	id, _ := strconv.ParseInt(key, 10, 64)
	dbconfig.Connector.Where("id = ?", id).Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
