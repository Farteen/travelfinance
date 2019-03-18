package event

import (
	"github.com/Farteen/travelfinance/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	eventGroup := engine.Group("/event")
	eventGroup.Use(middleware.CookieUIDAuthMiddleWare)
	{
		eventGroup.GET("/all", allEvents)
		eventGroup.POST("/new", eventCreation)
		eventGroup.GET("/{id}", eventItem)
		eventGroup.DELETE("/{id}", eventDeletion)
	}

}