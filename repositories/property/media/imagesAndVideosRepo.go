package media

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
)

func CreatePropertyImageOrVideo(propertyImagesOrVideos *[]propertyModels.PropertyImageOrVideo) bool {
	db.DB.Create(propertyImagesOrVideos)
	return true
}

func GetAllPropertiesImagesOrVideos() []propertyModels.PropertyImageOrVideo {
	var mediaList = []propertyModels.PropertyImageOrVideo{}
	err := db.DB.Find(&mediaList)
	if err != nil {
		println(err.Error, err.Name())
	}
	return mediaList
}

func GetPropertyImageOrVideoById(propertyMediaId string) *propertyModels.PropertyImageOrVideo {
	var media propertyModels.PropertyImageOrVideo
	result := db.DB.First(&media, propertyMediaId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &media
}

func GetAllPropertyImagesOrVideosByPropertyId(propertyId int) []propertyModels.PropertyImageOrVideo {
	var mediaList = []propertyModels.PropertyImageOrVideo{}
	err := db.DB.Where("property_id= ?", propertyId).Find(&mediaList)
	if err != nil {
		println(err.Error, err.Name())
	}
	return mediaList
}

func UpdatePropertyImageOrVideo(propertyImageOrVideoUpdate *propertyModels.PropertyImageOrVideo) bool {
	db.DB.Save(propertyImageOrVideoUpdate)
	return true
}

func DeletePropertyImageOrVideo(propertyImageOrVideoId string) bool {
	db.DB.Unscoped().Delete(&propertyModels.PropertyImageOrVideo{}, propertyImageOrVideoId)
	return true
}
