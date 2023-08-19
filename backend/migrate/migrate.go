package main

import (
	"moneyball/initializers"
	"moneyball/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Account{})
	initializers.DB.AutoMigrate(&models.User{})
}