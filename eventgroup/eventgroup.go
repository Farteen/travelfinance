package eventgroup

import (
	"context"
	"github.com/Farteen/travelfinance/cookie"
	"github.com/Farteen/travelfinance/event"
	"github.com/Farteen/travelfinance/mongoclient"
	"github.com/Farteen/travelfinance/response"
	"github.com/Farteen/travelfinance/util"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"net/http"
)

func allEventGroups(ctx *gin.Context) {
	userId, userIdCookieErr := ctx.Cookie(cookie.UserCookieUID)
	if userIdCookieErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {
		}{}))
		return
	}

	userIdOID, userIdOIDErr := primitive.ObjectIDFromHex(userId)
	if userIdOIDErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {
		}{}))
		return
	}

	eventGroup := EventGroup{}
	eventGroup.UserId = userIdOID
	eventGroupBindErr := ctx.Bind(&eventGroup)
	if eventGroupBindErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {
		}{}))
		return
	}

	findAllFilter := bson.M {"user_id": userIdOID}
	eventGroupCount, eventGroupCountErr := mongoclient.Collection(event.MongoEventCollection).
		Count(context.Background(),
			findAllFilter,
			options.Count())
	if eventGroupCountErr != nil || eventGroupCount >= EventGroupLimitCount {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "超出最大创建数量", struct {
		}{}))
		return
	}

	insertResult, insertResultErr := mongoclient.Collection(MongoDBEventGroupCollection).
		InsertOne(context.Background(), eventGroup, options.InsertOne())
	if insertResultErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {
		}{}))
		return
	}
	eventGroup.EventGroupId = util.MongoDBOID(insertResult)
	ctx.JSON(http.StatusOK, response.NewResponse(0, "创建成功", eventGroup))
}


func addEventToGroup(ctx *gin.Context) {
	eventGroupAdding := EventGroupAdding{}
	eventGroupAddingErr := ctx.Bind(&eventGroupAdding)
	if eventGroupAddingErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {
		}{}))
		return
	}

	eventGroupFilter := bson.M{"_id": eventGroupAdding.EventGroupId}
	findResult := mongoclient.Collection(MongoDBEventGroupCollection).
		FindOne(context.Background(),
			eventGroupFilter,
			options.FindOne())
	if findResult.Err() != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "更新失败", struct {
		}{}))
		return
	}
	eventGroupInDB := EventGroup{}
	eventGroupDecodingErr := findResult.Decode(&eventGroupInDB)
	if eventGroupDecodingErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "更新失败", struct {
		}{}))
		return
	}
	eventGroupInDB.EventItems = append(eventGroupInDB.EventItems, eventGroupAdding.EventId)
	_, updateEventGroupResultErr := mongoclient.Collection(MongoDBEventGroupCollection).
		UpdateOne(context.Background(),
			eventGroupInDB,
			options.Update())
	if updateEventGroupResultErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "更新失败", struct {
		}{}))
		return
	}
	ctx.JSON(http.StatusOK, response.NewResponse(0, "添加成功", eventGroupInDB))
}

func createEventGroup(ctx *gin.Context) {
	userId, userIdCookieErr := ctx.Cookie(cookie.UserCookieUID)
	if userIdCookieErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {
		}{}))
		return
	}

	userIdOID, userIdOIDErr := primitive.ObjectIDFromHex(userId)
	if userIdOIDErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {
		}{}))
		return
	}

	eventGroup := EventGroup{}
	eventGroup.UserId = userIdOID

	eventGroupAddingErr := ctx.Bind(&eventGroup)
	if eventGroupAddingErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {
		}{}))
		return
	}

	eventGroupResult, updateEventGroupResultErr := mongoclient.Collection(MongoDBEventGroupCollection).
		InsertOne(context.Background(),
		eventGroup,
			options.InsertOne())
	if updateEventGroupResultErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "更新失败", struct {
		}{}))
		return
	}
	eventGroup.EventGroupId = util.MongoDBOID(eventGroupResult)
	ctx.JSON(http.StatusOK, response.NewResponse(0, "添加成功", eventGroup))
}