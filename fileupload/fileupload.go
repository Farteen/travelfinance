package fileupload

import (
	"github.com/Farteen/travelfinance/response"
	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
)

func UploadFile(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Info(file.Filename)
	//OSS
	
	ctx.JSON(http.StatusOK, response.NewResponse(0, "upload file success", ))
}
