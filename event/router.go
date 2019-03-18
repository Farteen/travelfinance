package event

import (
	"github.com/Farteen/travelfinance/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	eventRouter := engine.Group("/event")
	eventRouter.Use(middleware.CookieUIDAuthMiddleWare)
	{
		eventRouter.GET("/all", allEvents)
		eventRouter.POST("/new", eventCreation)
		eventRouter.GET("/{id}", eventItem)
		eventRouter.DELETE("/{id}", eventDeletion)
	}



}