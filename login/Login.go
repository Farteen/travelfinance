package login

import (
	"context"
	"encoding/hex"
	"github.com/Farteen/travelfinance/mongoclient"
	"github.com/Farteen/travelfinance/redisclient"
	"github.com/Farteen/travelfinance/response"
	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/mongo/driver/uuid"
	"time"
)

const (
	UserRegisterNoError = 0
	UserRegisterInputError  = iota + 1000
	UserRegisterUserNameError
	UserRegisterUserPhoneError
	UserRegisterIdGenerationError
	UserRegisterPersistanceError
)

const (
	UserRegisterInputErrorMsg = "用户输入信息有误"
	UserRegisterUserNameErrorMsg = "用户已存在"
	UserRegisterUserPhoneErrorMsg = "用户已存在"
	UserRegisterIdGenerationErrorMsg = "用户ID生成错误"
	UserRegisterPersistanceErrorMsg = "用户信息存储问题"
	UserRegisterNoErrorMsg = "用户注册成功"
)

const (
	UserIdRedisSetKey = "global:user:all"
)

const (
	MongoDBUserCollection = "users"
)

type UserRegister struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone"`
}

func userRegister(ctx *gin.Context) {
	userRegister := UserRegister{}
	ctx.Bind(&userRegister)
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
	//TODO: unit testing
	userId := hex.EncodeToString(uuid[0:])
	userNameKey := "user:" + "username:" + userRegister.UserName
	userPhoneToId := "user:" + "userphone:" + userRegister.PhoneNumber
	userNameResult, userNameError := redisclient.RedisClient().SetNX(userNameKey, userId, time.Duration(0)).Result()
	if userNameError != nil || userNameResult == false {
		ctx.JSON(200, response.NewResponse(UserRegisterUserNameError, UserRegisterUserNameErrorMsg, struct {}{}))
		return
	}
	userPhoneResult, userPhoneError := redisclient.RedisClient().SetNX(userPhoneToId, userId, time.Duration(0)).Result()
	if userPhoneError != nil || userPhoneResult == false {
		ctx.JSON(200, response.NewResponse(UserRegisterUserPhoneError, UserRegisterUserPhoneErrorMsg, struct {}{}))
		return
	}
	userIdResult, userIdError := redisclient.RedisClient().SetNX(UserIdRedisSetKey, userId, time.Duration(0)).Result()
	if userIdError != nil || userIdResult == false {
		ctx.JSON(200,
			response.NewResponse(
				UserRegisterIdGenerationError,
				UserRegisterIdGenerationErrorMsg,
				struct {}{}))
		return
	}
	mongodbInsertResult, mongodbInsertErr := mongoclient.MongoDBCollectionWithName(MongoDBUserCollection).InsertOne(
		context.Background(),
		userRegister,
		options.InsertOne())
	insertId := mongodbInsertResult.InsertedID
	if _, ok := insertId.(primitive.ObjectID); !ok || mongodbInsertErr != nil {
		ctx.JSON(200, response.NewResponse(UserRegisterPersistanceError, UserRegisterPersistanceErrorMsg, struct {
		}{}))
		return
	}
	ctx.JSON(200, response.NewResponse(UserRegisterNoError, UserRegisterNoErrorMsg, struct {}{}))
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
