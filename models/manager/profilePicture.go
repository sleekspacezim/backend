package manager

import baseModel "SleekSpace/models"

type ManagerProfilePicture struct {
	baseModel.MyModel
	Id          int    `json:"id" gorm:"primary_key"`
	ManagerId   int    `json:"managerId"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}
