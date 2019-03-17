package cookie


const (
	UserCookieUID = "com.travelfinance.glassed.uid"
	CookieMaxAge = 60 * 60 * 24 * 30
)

const (
	UserCookieUIDNotFoundErr = iota + 1000
)

const (
	UserCookieUIDNotFoundErrMsg = "信息不完整"
)