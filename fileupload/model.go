package fileupload

type UploadFileItem struct {
	FileURL string	`json:"file_url" bson:"file_url"`
	FileName string `json:"file_name" bson:"file_name"`
}


