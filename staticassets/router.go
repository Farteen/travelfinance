package staticassets

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouter(engine *gin.Engine) {
	//engine.StaticFile("/static", "/Users/glassesd/Desktop/static_assets")
	engine.Static("/static", "./static_assets")
	engine.StaticFS("/sfs", http.Dir("static_assets_fs"))
}
