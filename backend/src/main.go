package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", home)
	r.Run()
}

func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "Hello",
	})
}
