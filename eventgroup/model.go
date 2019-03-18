package eventgroup

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type EventGroup struct {
	EventGroupId primitive.ObjectID `bson:"_id,omitempty" json:"event_group_id"`
	UserId primitive.ObjectID `bson:"user_id" json:"user_id"`
	EventItems []primitive.ObjectID `bson:"event_items" json:"event_items"`
}

type EventGroupAdding struct {
	EventGroupId primitive.ObjectID `bson:"_id,omitempty" json:"event_group_id"`
	EventId primitive.ObjectID `bson:"event_id" json:"event_id"`
}