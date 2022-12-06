package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"final_project_3/database/models"
)

var (
	host   = os.Getenv("DB_HOST")
	port   = os.Getenv("DB_PORT")
	user   = os.Getenv("DB_USER")
	pass   = os.Getenv("DB_PASS")
	dbname = os.Getenv("DB_NAME")
)

func ConnectionDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{})
	if err != nil {
		return db, err
	}

	return db, nil
}

// func ConnectionDB() (*gorm.DB, error) {
// 	db, err := gorm.Open(sqlite.Open("./kanban.db"), &gorm.Config{})
// 	if err != nil {
// 		return db, err
// 	}

// 	err = db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{})
// 	if err != nil {
// 		return db, err
// 	}

// 	return db, nil
// }
