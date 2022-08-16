package database

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
)


var Connector *gorm.DB


type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}


func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}



var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}