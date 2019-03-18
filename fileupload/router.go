package fileupload

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engin *gin.Engine) {
	group := engin.Group("/upload")
	//group.Use(middleware.RedisSessionMiddleWare)
	{
		group.POST("/file", UploadFile)
	}
}