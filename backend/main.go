package main

import (
	"moneyball/controllers"
	"moneyball/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/account", controllers.CreateAccount)
	r.GET("/accounts", controllers.IndexAccounts)
	r.GET("/account/:id", controllers.ShowAccount)
	r.PUT("/account/:id", controllers.UpdateAccount)
	r.DELETE("/account/:id", controllers.DeleteAccount)

	r.POST("/user", controllers.CreateUser)
	r.GET("/users", controllers.IndexUsers)
	r.GET("/user/:id", controllers.ShowUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.Run()
}