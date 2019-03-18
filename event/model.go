package event

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type EventItem struct {
	EventID primitive.ObjectID `bson:"_id,omitempty" json:"eventId"`
	CreatorID primitive.ObjectID	`bson:"creatorId" json:"creatorId"`
	Name string	`bson:"name" json:"name"`
	Description string	`bson:"desc" json:"desc"`
	Images []string	`bson:"images" json:"images"`
}



