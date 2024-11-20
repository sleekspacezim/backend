package media

type PropertyImageOrVideoCreationWithNoPropertyIdDto struct {
	File        string `json:"file"`
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}

type PropertyImageOrVideoCreationWithPropertyIdDto struct {
	PropertyId   int    `json:"propertyId"`
	File         string `json:"file"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
	ContentType  string `json:"contentType"`
	FileType     string `json:"fileType"`
	PropertyType string `json:"propertyType"`
}
type PropertyImagesAndVideosDto struct {
	Files []PropertyImageOrVideoCreationWithPropertyIdDto `json:"files"`
}

type PropertyImageOrVideoUpdateDto struct {
	Id           int    `json:"id"`
	PropertyId   int    `json:"propertyId"`
	File         string `json:"file"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
	ContentType  string `json:"contentType"`
	FileType     string `json:"fileType"`
	PropertyType string `json:"propertyType"`
}

type PropertyImageOrVideoResponseDto struct {
	Id           int    `json:"id"`
	PropertyId   int    `json:"propertyId"`
	Uri          string `json:"uri"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
	ContentType  string `json:"contentType"`
	FileType     string `json:"fileType"`
	PropertyType string `json:"propertyType"`
}
