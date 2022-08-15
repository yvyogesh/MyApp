package controllers

import (
	"MyApp/backend/m/initializers"
	"MyApp/backend/m/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := models.User{}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusBadRequest, user)
}

func GetUsers(c *gin.Context) {

}
