package login

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type UserInfo struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	PhoneNumber string `json:"phone" bson:"phone"`
	Token string `json:"token"`
}

type JwtToken struct {
	Token string `json:"token"`
}