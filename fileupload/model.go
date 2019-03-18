package fileupload

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type UploadFileItem struct {
	FileURL string	`json:"file_url" bson:"file_url"`
	FileName string `json:"file_name" bson:"file_name"`
}

type UploadFileDBEntity struct {
	*UploadFileItem
	FileID primitive.ObjectID `json:"file_id" bson:"_id"`
}

