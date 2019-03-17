package fileupload

import (
	"github.com/Farteen/travelfinance/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engin *gin.Engine) {
	engin.Group("/upload")
	engin.Use(middleware.RedisSessionMiddleWare)
	{
		engin.POST("", UploadFile)
	}
}