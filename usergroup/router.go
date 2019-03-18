package usergroup

import (
	"github.com/Farteen/travelfinance/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	group := engine.Group("/group")
	group.Use(middleware.CookieUIDAuthMiddleWare)
	{
		group.GET("/all", UserGroupList)
		group.POST("/new", UserGroupCreation)
		group.POST("/{id}", UserGroupAddUser)
		group.DELETE("/{id}", UserGroupDeletion)
	}
}