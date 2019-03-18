package usergroup

import (
	"context"
	"github.com/Farteen/travelfinance/mongoclient"
	"github.com/Farteen/travelfinance/response"
	"github.com/Farteen/travelfinance/util"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"net/http"
)

func UserGroupCreation(ctx *gin.Context) {
	userGroup := UserGroupRequest{}
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

}

func UserGroupDeletion(ctx *gin.Context) {

}

func UserGroupAddUser(ctx *gin.Context) {

}