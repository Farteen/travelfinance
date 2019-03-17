package login

type UserUnloggedinInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone"`
	Token string `json:"token"`
	ID string `json:"_id"`
}

type JwtToken struct {
	Token string `json:"token"`
}