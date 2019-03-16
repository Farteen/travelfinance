package login

import (
	"encoding/binary"
	"github.com/Farteen/travelfinance/redisclient"
	"github.com/Farteen/travelfinance/response"
	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/mongodb/mongo-go-driver/x/mongo/driver/uuid"
)

const (
	UserRegisterInputError  = iota + 1000
)

const (
	UserRegisterInputErrorMsg = "用户输入信息有误"
)

type UserRegister struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone"`
}

func userRegister(ctx *gin.Context) {
	userRegister := UserRegister{}
	ctx.BindJSON(userRegister)
	if len(userRegister.UserName) == 0 ||
		len(userRegister.Password) == 0 ||
		len(userRegister.PhoneNumber) == 0 {
		ctx.JSON(200, response.NewResponse(UserRegisterInputError, UserRegisterInputErrorMsg, struct {}{}))
		return
	}
	//TODO: user info validation, like username wihout blackspace

	//TODO: user id generator
	uuid, uuidErr := uuid.New()
	if uuidErr != nil {
		log.Fatal("uuid error")
		return
	}
	userId := string(uuid[0:])
	userKey := "user:" + "username:" + userRegister.UserName + ":userId:" + userId

	redisclient.RedisClient().HSet()
}

func userLogin(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{
		"key": "user login",
	})
}

func RegisterLoginRouter(engine *gin.Engine) {
	register := engine.Group("/register")
	{
		register.POST("", userRegister)
	}

	login := engine.Group("/login")
	{
		login.POST("", userLogin)
	}
}
