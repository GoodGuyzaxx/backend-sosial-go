package database

import (
	"fmt"
	"log"
	"os"

	"zaxx/backend/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)


	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass,dbHost,dbPort,dbName)

	var err error
	DB , err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatal("test")
		log.Fatal("Failed To Connect to database:", err)
	}

	fmt.Println("Database Conntected")

	err  = DB.AutoMigrate(&model.Post{}, &model.User{})
	if err != nil {
		log.Fatal("Failed To Migrate: ", err)
	}

	fmt.Println("Database Migrate Success")


}