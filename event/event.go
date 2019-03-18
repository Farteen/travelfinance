package event

import (
	"context"
	"github.com/Farteen/travelfinance/cookie"
	"github.com/Farteen/travelfinance/mongoclient"
	"github.com/Farteen/travelfinance/response"
	"github.com/Farteen/travelfinance/util"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"net/http"
)

func allEvents(ctx *gin.Context) {
	userId, err := ctx.Cookie(cookie.UserCookieUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, struct {}{})
		return
	}
	filter := bson.M {
		"userId": userId,
	}
	cursor, err := mongoclient.Collection(MongoEventCollection).Find(
		context.Background(),
		filter,
		options.Find().SetBatchSize(EventMongoQueryBatchSize))
	eiList := make([]EventItem, 0)
	for cursor.Next(context.Background()) {
	 	ei := EventItem{}
	 	eiErr := cursor.Decode(ei)
		if eiErr == nil {
			eiList = append(eiList, ei)
		}
	}
	ctx.JSON(http.StatusOK, response.NewResponse(0, "event list success", eiList))
}

func eventCreation(ctx *gin.Context) {
	ei := EventItem{}
 	eiErr := ctx.Bind(&ei)
 	if eiErr != nil ||
 		len(ei.GroupID.Hex()) == 0 ||
 		len(ei.Name) == 0 ||
 		len(ei.CreatorID.Hex()) == 0 ||
 		len(ei.Images) == 0 {
 		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "bad request", struct {}{}))
		return
	}

 	insertResult, err := mongoclient.Collection(MongoEventCollection).
 		InsertOne(
 			context.Background(),
 			ei,
 			options.InsertOne())

	if err != nil {
		//TODO: insert mongodb error
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "bad request", struct {}{}))
		return
	}
 	ei.EventID = util.MongoDBOID(insertResult)
	ctx.JSON(http.StatusOK, response.NewResponse(0, "event ok", ei))
}

func eventItem(ctx *gin.Context) {

}

func eventDeletion(ctx *gin.Context) {

}
