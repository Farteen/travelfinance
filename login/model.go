package login

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type UserUnloggedinInfo struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone"`
	Token string `json:"token"`
}

type JwtToken struct {
	Token string `json:"token"`
}