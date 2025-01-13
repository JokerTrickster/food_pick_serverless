package entity

import "mime/multipart"

type ImageUploadFoodEntity struct {
	Image *multipart.FileHeader `json:"image"`
}
