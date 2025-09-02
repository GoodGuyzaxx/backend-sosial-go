package model

import(
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConncetDatabase(){
	dsn := "root:zaxx@tcp(127.0.01:3306)/db_go_api"
	database , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Fail to Connect")
	}

	database.AutoMigrate(&Post{})

	DB = database
}