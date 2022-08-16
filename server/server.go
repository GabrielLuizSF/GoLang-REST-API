package server

import (
	"log"
	"net/http"
	"user_api-golang/database"
	"user_api-golang/models"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "password364824672483472463264873264782648",
			DB:         "users",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&models.User{})
}

func StartServer(){
	initDB()
	log.Println("Starting the HTTP server on port 5700")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":5700", router))
}