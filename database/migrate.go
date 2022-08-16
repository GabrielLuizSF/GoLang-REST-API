package database

import (
	"log"
	"user_api-golang/models"
)


func Migrate(table *models.User) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}