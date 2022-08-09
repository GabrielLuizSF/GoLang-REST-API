package migrations

import (
	"github.com/GabrielLuizSF/GoLang-REST-API/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}