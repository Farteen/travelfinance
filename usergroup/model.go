package usergroup

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type UserGroupRequest struct {
	GroupName string `json:"group_name" bson:"group_name"`
	MaintainerId primitive.ObjectID `json:"maintainer_id" bson:"maintainer_id"`
	Users []primitive.ObjectID `json:"users" bson:"users"`
}

type UserGroupDBEntity struct {
	*UserGroupRequest
	UserGroupId primitive.ObjectID	`bson:"_id"`
}


