package eventgroup

import (
	"github.com/Farteen/travelfinance/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	eventGroupRouter := engine.Group("/eventgroup")
	eventGroupRouter.Use(middleware.CookieUIDAuthMiddleWare)
	{
		eventGroupRouter.GET("/all", allEventGroups)
		eventGroupRouter.POST("/new", createEventGroup)
		eventGroupRouter.POST("/id/:id", addEventToGroup)
	}
}