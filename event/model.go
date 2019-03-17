package event

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type EventItem struct {
	EventID *primitive.ObjectID `bson:"_id" json:"eventId"`
	Creator string	`bson:"creator" json:"creator"`
	Name string	`bson:"name" json:"name"`
	Description string	`bson:"desc" json:"desc"`
	Images []string	`bson:"images" json:"images"`
	EventGroupId string `bson:"event_groupId" json:"event_groupId"`
}


