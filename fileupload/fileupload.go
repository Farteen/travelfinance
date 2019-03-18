package fileupload

import (
	"context"
	"github.com/Farteen/travelfinance/mongoclient"
	"github.com/Farteen/travelfinance/response"
	"github.com/Farteen/travelfinance/staticassets"
	"github.com/Farteen/travelfinance/util"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"net/http"
)

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(
			FileUploadMongoDBInsertionErr,
			FileUploadMongoDBInsertionErrMsg,
			struct {}{}))
		return
	}
	//OSS upload and then write to db

	fileItem := UploadFileItem{}
	fileItem.FileName = file.Filename
	//TODO:insert item to database
	insertOID, insertErr := mongoclient.Collection(MongoDBStaticFileCollection).InsertOne(
		context.Background(),
		fileItem,
		options.InsertOne())
	if insertErr != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(
			FileUploadMongoDBInsertionErr,
			FileUploadMongoDBInsertionErrMsg,
			struct {}{}))
		return
	}

	fileName := util.MongoDBHexID(insertOID)
	//fileType := file.Header.Get("Content-Type")
	staticAssetServer := StaticAssetsHostName
	localFilePath := staticassets.StaticAssetRelativeLocalPathComponent + "/" + fileName + ".png"
	fileURL := staticassets.StaticAssetRelativeURLPathComponent + "/" + fileName + ".png"
	//TODO:save on local system
	fileSaveErr := ctx.SaveUploadedFile(file, localFilePath)
	if fileSaveErr != nil {
		ctx.JSON(http.StatusOK,
			response.NewResponse(FileUploadMongoDBInsertionErr,
				FileUploadMongoDBInsertionErrMsg,
				struct {}{}))
		return
	}

	fileItem.FileURL = staticAssetServer + fileURL

	ctx.JSON(http.StatusOK, response.NewResponse(0, "upload file success", fileItem))
}
