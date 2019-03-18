package main

import (
	"github.com/Farteen/travelfinance/event"
	"github.com/Farteen/travelfinance/fileupload"
	"github.com/Farteen/travelfinance/login"
	"github.com/Farteen/travelfinance/staticassets"
	"github.com/Farteen/travelfinance/usergroup"
	"github.com/gin-gonic/gin"
)


type RegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	PasswordAgain string `json:"password_again"`
}


func main() {
	router := gin.Default()
	login.RegisterRouter(router)
	event.RegisterRouter(router)
	fileupload.RegisterRouter(router)
	usergroup.RegisterRouter(router)
	staticassets.RegisterRouter(router)
	router.Run("localhost:8080")
}


