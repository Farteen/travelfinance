package staticassets

import "github.com/gin-gonic/gin"

func RegisterRouter(engine *gin.Engine) {
	engine.StaticFile("/static", "./")
}
