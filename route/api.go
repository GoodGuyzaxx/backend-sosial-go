package route

import(
	"os"

	"zaxx/backend/controller"
	"github.com/gin-gonic/gin"
	"zaxx/backend/config"
)

func RouteMain(){
	router := gin.Default()
	config.LoadEnv()
		
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "Hello World",
		})
	})

	router.GET("/api/posts", controller.FindPost)
	router.POST("/api/posts", controller.StorePost)
	router.GET("/api/posts/:id", controller.FindPostById)
	router.PUT("/api/posts/:id", controller.UpdatePost)
	router.DELETE("/api/posts/:id", controller.DeletePost)

	router.Run(":"+ os.Getenv("APP_PORT"))

}