package login

const (
	UserRegisterNoError = 0
	UserRegisterInputError  = iota + 1000
	UserRegisterUserNameError
	UserRegisterUserPhoneError
	UserRegisterIdGenerationError
	UserRegisterPersistanceError
)

const (
	UserLoginNoError = 0
	UserLoginInputError = iota + 1000
	UserLoginUserNotFoundError
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
	UserLoginNoErrorMsg = "用户登录成功"
	UserLoginInputErrorMsg = "用户信息输入有误"
	UserLoginUserNotFoundErrorMsg = "用户不存在,或密码错误"
)

const (
	UserIdsRedisSetKey = "global:userids"
	UserNamesRedisSetKey = "global:usernames"
	UserPhonesRedisSetKey = "global:userphones"
)

const (
	MongoDBUserCollection = "users"
)
