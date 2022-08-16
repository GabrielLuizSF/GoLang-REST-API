package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"user_api-golang/database"
	"user_api-golang/models"

	"github.com/gorilla/mux"
)

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var Users []models.User
	database.Connector.Find(&Users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var User models.User
	database.Connector.First(&User, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var User models.User
	json.Unmarshal(requestBody, &User)

	database.Connector.Create(User)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(User)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var User models.User
	json.Unmarshal(requestBody, &User)
	database.Connector.Save(&User)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(User)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var user models.User
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
