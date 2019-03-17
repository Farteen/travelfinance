package login

import "github.com/gin-gonic/gin"

func RegisterRouter(engine *gin.Engine) {
	engine.POST("/register", userRegister)
	engine.POST("/login", userLogin)
	engine.GET("/cookies", readCookie)
}
