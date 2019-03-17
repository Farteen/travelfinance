package event

type EventItem struct {
	EventID string `bson:"_id" json:"eventId"`
	Creator string	`bson:"creator" json:"creator"`
	Name string	`bson:"name" json:"name"`
	Description string	`bson:"desc" json:"desc"`
	Images []string	`bson:"images" json:"images"`
	EventGroupId string `bson:"event_groupId" json:"event_groupId"`
}


