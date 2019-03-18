package login

import (
	"context"
	"github.com/Farteen/travelfinance/cookie"
	"github.com/Farteen/travelfinance/mongoclient"
	"github.com/Farteen/travelfinance/redisclient"
	"github.com/Farteen/travelfinance/response"
	"github.com/Farteen/travelfinance/util"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"net/http"
)

func userRegister(ctx *gin.Context) {
	userRegister := UserInfo{}
	err := ctx.Bind(&userRegister)
	if err != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(
			UserRegisterInputError,
			UserRegisterInputErrorMsg,
			struct {}{}))
		return
	}

	if len(userRegister.UserName) == 0 ||
		len(userRegister.Password) == 0 ||
		len(userRegister.PhoneNumber) == 0 {
		ctx.JSON(http.StatusOK, response.NewResponse(UserRegisterInputError, UserRegisterInputErrorMsg, struct {}{}))
		return
	}
	//TODO: user info validation,
	// like username wihout blackspace,
	// like phone validation, etc.

	//userNameKey := "user:" + "username:" + userRegister.UserName
	userNameExists, userNameInRedisErr := redisclient.RedisClient().SIsMember(
		UserNamesRedisSetKey,
		userRegister.UserName).Result()
	if userNameInRedisErr != nil || userNameExists {
		ctx.JSON(http.StatusOK, response.NewResponse(
			UserRegisterUserNameError,
			UserRegisterUserNameErrorMsg,
			struct {}{}))
		return
	}

	userPhoneExists, userPhoneInRedisErr := redisclient.RedisClient().SIsMember(
		UserPhonesRedisSetKey,
		userRegister.PhoneNumber).Result()
	if userPhoneInRedisErr != nil || userPhoneExists {
		ctx.JSON(http.StatusOK, response.NewResponse(
			UserRegisterUserNameError,
			UserRegisterUserNameErrorMsg,
			struct {}{}))
		return
	}

	//TODO:user password encryption
	userRegister.Password = util.MD5String(userRegister.Password)

	//TODO: user id generator
	mongodbInsertResult, mongodbInsertErr := mongoclient.Collection(MongoDBUserCollection).
		InsertOne(
		context.Background(),
		userRegister,
		options.InsertOne())
	insertId := mongodbInsertResult.InsertedID
	_, ok := insertId.(primitive.ObjectID)
	if !ok || mongodbInsertErr != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(
			UserRegisterPersistanceError,
			UserRegisterPersistanceErrorMsg,
			struct {}{}))
		return
	}

	redisclient.RedisClient().SAdd(UserNamesRedisSetKey, userRegister.UserName)
	redisclient.RedisClient().SAdd(UserPhonesRedisSetKey, userRegister.PhoneNumber)

	ctx.JSON(http.StatusOK, response.NewResponse(UserRegisterNoError, UserRegisterNoErrorMsg, struct {}{}))
}

func userLogin(ctx *gin.Context) {

	userLogin := UserInfo{}
	err := ctx.Bind(&userLogin)
	if err != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(UserLoginInputError, UserLoginInputErrorMsg, struct {}{}))
		return
	}
	filter := bson.M{"phonenumber" : userLogin.PhoneNumber}
	singleResult := mongoclient.Collection(MongoDBUserCollection).
		FindOne(context.Background(),
			filter,
			options.FindOne())

	if singleResult.Err() != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(
			UserLoginUserNotFoundError,
			UserLoginInputErrorMsg, struct {}{}))
		return
	}

	loginPwdMd5Str := util.MD5String(userLogin.Password)
	//MARK:Decode方法带的是一个指针
	decodeUserErr := singleResult.Decode(&userLogin)
	if decodeUserErr != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(
			UserLoginUserNotFoundError,
			UserLoginUserNotFoundErrorMsg,
			struct {}{}))
		return
	}
	if loginPwdMd5Str != userLogin.Password {
		ctx.JSON(http.StatusOK, response.NewResponse(
			UserLoginUserNotFoundError,
			UserLoginUserNotFoundErrorMsg,
			struct {}{}))
		return
	}

	ctx.SetCookie(
		cookie.UserCookieUID,
		userLogin.ID.Hex(),
		cookie.CookieMaxAge,
		"/",
		"localhost",
		false,
		true)
	ctx.JSON(
		http.StatusOK,
		response.NewResponse(
			UserLoginNoError,
			UserLoginNoErrorMsg,
			struct {}{}))
}

func readCookie(ctx *gin.Context) {
	result, err := ctx.Cookie(cookie.UserCookieUID)
	if err != nil {
		ctx.JSON(http.StatusOK, response.NewResponse(-1, "cookie error", struct {}{}))
		return
	}
	ctx.JSON(http.StatusOK, response.NewResponse(0, "cookie read success", result))
}
