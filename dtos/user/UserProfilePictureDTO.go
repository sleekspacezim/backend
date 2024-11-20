package dtos

type UserProfilePictureResponseDTO struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}

type UserProfilePictureUpdateDTO struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}
type UserProfilePictureCreationDTO struct {
	UserId      int    `json:"userId"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}
