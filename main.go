package main

import (
	"github.com/Farteen/travelfinance/login"
	"github.com/gin-gonic/gin"
)


type RegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	PasswordAgain string `json:"password_again"`
}


func main() {
	router := gin.Default()
	login.RegisterLoginRouter(router)

	router.Run("localhost:8080")
}


