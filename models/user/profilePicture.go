package models

import baseModel "SleekSpace/models"

type UserProfilePicture struct {
	baseModel.MyModel
	Id          int    `json:"id" gorm:"primary_key"`
	UserId      int    `json:"userId"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}
