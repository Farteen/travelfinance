package event

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type EventItem struct {
	EventID primitive.ObjectID `bson:"_id,omitempty" json:"eventId"`
	CreatorID primitive.ObjectID	`bson:"creatorId" json:"creatorId"`
	Users []primitive.ObjectID `bson:"users" json:"users"`
	Name string	`bson:"name" json:"name"`
	EventType int `bson:"event_type" json:"event_type"`
	CreateTimestamp int32 `bson:"create_timestamp" json:"create_timestamp"`
	Description string	`bson:"desc" json:"desc"`
	Images []string	`bson:"images" json:"images"`
}

