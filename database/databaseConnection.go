package database

import (
	authConfig "github.com/codertjay/golang-gorm-allauth/config"
	"github.com/codertjay/golang-gorm-allauth/models"
	"gorm.io/gorm"
	"log"
)

func Instance() (*gorm.DB, error) {
	/* This get the config instance */
	config := authConfig.AllauthConfig{}
	/* Get the database which was set on the allauth config by the user*/
	database := config.GetAllauthConfig().Database
	/* AutoMigrate the database if it has any changes*/
	err := database.AutoMigrate(&models.User{})
	if err != nil {
		log.Panicln("Failed to migrate database", err)
		return nil, err
	}
	return database, nil
}

var Client, _ = Instance()
