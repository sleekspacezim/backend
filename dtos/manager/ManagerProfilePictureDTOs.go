package dtos

type ManagerProfilePictureCreationDTO struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}

type ManagerProfilePictureUpdateDTO struct {
	Id          int    `json:"id"`
	ManagerId   int    `json:"managerId"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}

type ManagerProfilePictureResponseDTO struct {
	Id          int    `json:"id"`
	ManagerId   int    `json:"managerId"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}
