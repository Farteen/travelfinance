package event

import (
	"context"
	"github.com/Farteen/travelfinance/cookie"
	"github.com/Farteen/travelfinance/mongoclient"
	"github.com/Farteen/travelfinance/response"
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
		options.Find().SetBatchSize(10))
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

func eventItem(ctx *gin.Context) {

}

func eventCreation(ctx *gin.Context) {

}

func eventDeletion(ctx *gin.Context) {

}
