package main

import (
	"MyApp/backend/m/controllers"
	"MyApp/backend/m/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVeriables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", home)
	r.GET("/users", controllers.GetUsers)
	r.POST("/user", controllers.CreateUser)
	r.Run()
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello",
	})
}
