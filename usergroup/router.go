package usergroup

import (
	"github.com/Farteen/travelfinance/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/group")
	group.Use(middleware.CookieUIDAuthMiddleWare)
	{
		group.GET("", UserGroupList)
		group.POST("/new", UserGroupCreation)
		group.DELETE("/:id", UserGroupDeletion)
		group.POST("/:id/:uid", UserGroupAddUser)
	}
}