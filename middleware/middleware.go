package middleware

import (
	"github.com/Farteen/travelfinance/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CookieUIDAuthMiddleWare(ctx *gin.Context) {
	result, err := ctx.Cookie(cookie.UserCookieUID)
	if len(result) == 0 || err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, struct {}{})
		return
	}
	ctx.Next()
}