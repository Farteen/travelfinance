package event

import (
	"github.com/Farteen/travelfinance/cookie"
	"github.com/gin-gonic/gin"
)

type AllEventsRequest struct {

}

func allEvents(ctx *gin.Context) {
	ctx.Cookie(cookie.UserCookieUID)
}

func eventItem(ctx *gin.Context) {

}

func eventCreation(ctx *gin.Context) {

}

func eventDeletion(ctx *gin.Context) {

}
