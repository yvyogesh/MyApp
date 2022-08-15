package main

import (
	"MyApp/backend/m/initializers"
	"MyApp/backend/m/models"
)

func init() {
	initializers.LoadEnvVeriables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
