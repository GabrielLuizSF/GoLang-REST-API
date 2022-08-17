package database

import "user_api-golang/models"

func InitDB() {
	config :=
		Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "password364824672483472463264873264782648",
			DB:         "users",
		}

	connectionString := GetConnectionString(config)
	err := Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	Migrate(&models.User{})
}
