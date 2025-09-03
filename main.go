package main

import (
	// "github.com/gin-gonic/gin"

	// "zaxx/backend/model"
	// "fmt"
	// "os"
	"zaxx/backend/config"
	"zaxx/backend/database"
	"zaxx/backend/route"
)

func main(){
	config.LoadEnv()

database.InitDB()
	route.RouteMain()

	
}