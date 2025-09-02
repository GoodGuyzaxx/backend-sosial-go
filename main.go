package main
 

import(
	// "github.com/gin-gonic/gin"

	"zaxx/backend/model"
	"zaxx/backend/route"
)

func main(){
	// router := gin.Default()

	model.ConncetDatabase()
	route.RouteMain()

	// router.Run(":3000")
}