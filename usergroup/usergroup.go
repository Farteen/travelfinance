package usergroup

import (
	"context"
	"github.com/Farteen/travelfinance/cookie"
	"github.com/Farteen/travelfinance/mongoclient"
	"github.com/Farteen/travelfinance/response"
	"github.com/Farteen/travelfinance/util"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"net/http"
)

func UserGroupCreation(ctx *gin.Context) {
	userId, cookieErr := ctx.Cookie(cookie.UserCookieUID)
	if cookieErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {}{}))
		return
	}
	userOIDStr, userOIDStrErr := primitive.ObjectIDFromHex(userId)
	if userOIDStrErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {}{}))
		return
	}

	userGroup := UserGroupRequest{}
	userGroup.MaintainerId = userOIDStr
	userGroupErr := ctx.Bind(&userGroup)
	if userGroupErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "请求错误", struct {}{}))
		return
	}
	findCountFilter := bson.M{"maintainer_id" : userGroup.MaintainerId}
	userHasCreatedGroupCount, userHasCreatedGroupCountErr := mongoclient.Collection(MongoDBUserGroupCollection).
		Count(context.Background(),
			findCountFilter,
			options.Count())
	if userHasCreatedGroupCountErr != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(1000, "MongoDB Count Error", struct {
		}{}))
		return
	}

	if userHasCreatedGroupCount >= UserGroupMaxCountLimit {
		ctx.JSON(http.StatusOK, response.NewResponse(1000, "Count out of limit", struct {
		}{}))
		return
	}

	insertResult, insertErr := mongoclient.Collection(MongoDBUserGroupCollection).
		InsertOne(context.Background(),
		userGroup,
		options.InsertOne())
	if insertErr != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(1000, "持久化失败", struct {
		}{}))
		return
	}

	userGroup.GroupID = util.MongoDBOID(insertResult)
	ctx.JSON(http.StatusOK, response.NewResponse(0, "创建成功", userGroup))
}

func UserGroupList(ctx *gin.Context) {
	userId, cookieErr := ctx.Cookie(cookie.UserCookieUID)
	if cookieErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000, "暂无权限", struct {}{}))
		return
	}
	userOID, userOIDErr := primitive.ObjectIDFromHex(userId)
	if userOIDErr != nil {
		ctx.JSON(http.StatusBadRequest, response.NewResponse(1000,"暂无权限", struct {}{}))
		return
	}
	findFilter := bson.M{"maintainer_id": userOID}
	findUserGroups, findErr := mongoclient.Collection(MongoDBUserGroupCollection).
		Find(context.Background(),
			findFilter,
			options.Find().SetBatchSize(UserGroupFindBatchSize))
	if findErr != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(1000, "查询失败", struct {}{}))
		return
	}
	listResult := make([]UserGroupRequest, 0)
	for findUserGroups.Next(context.Background()) {
		userGroup := UserGroupRequest{}
		decodeErr := findUserGroups.Decode(&userGroup)
		if decodeErr == nil {
			listResult = append(listResult, userGroup)
		}
	}
	ctx.JSON(http.StatusOK, response.NewResponse(0, "查询成功", listResult))
}

func UserGroupDeletion(ctx *gin.Context) {

}

func UserGroupAddUser(ctx *gin.Context) {

}