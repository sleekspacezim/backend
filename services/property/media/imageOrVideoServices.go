package media

import (
	"net/http"

	imagesOrVideosDtos "SleekSpace/dtos/property/media"
	propertyModels "SleekSpace/models/property"
	imagesOrVideosRepo "SleekSpace/repositories/property/media"
	generalServices "SleekSpace/services/property"
	"SleekSpace/storage"
	constants "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreatePropertyImageOrVideoWithPropertyId(c *gin.Context) {
	var files imagesOrVideosDtos.PropertyImagesAndVideosDto
	validateModelFields := validator.New()
	c.BindJSON(&files)

	modelFieldsValidationError := validateModelFields.Struct(files)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	if len(files.Files) > constants.ImagesOrVideosLimitPerProperty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can only upload " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty) + " images/videos"})
		return
	}

	propertyMediaFiles := imagesOrVideosRepo.GetAllPropertyImagesOrVideosByPropertyId(files.Files[0].PropertyId)
	if len(propertyMediaFiles) >= constants.ImagesOrVideosLimitPerProperty {
		c.JSON(http.StatusBadRequest, gin.H{"error": "images/videos limit of " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty) + " has been reached"})
		return
	}

	if len(propertyMediaFiles)+len(files.Files) > constants.ImagesOrVideosLimitPerProperty {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "images/videos limit of " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty) + " has been reached, you can only add " + generalUtilities.ConvertIntToString(constants.ImagesOrVideosLimitPerProperty-len(propertyMediaFiles)) + " images or videos without deleting some."})
		return
	}

	mediaList := propertyUtilities.MediaListWithPropertyId(files.Files)
	mediaUrls := storage.UploadFiles(mediaList, c)

	processedFiles := propertyUtilities.ConvertPropertyImagesOrVideosWithPropertyIdToModel(files.Files, mediaUrls)

	isImageOrVideoCreated := imagesOrVideosRepo.CreatePropertyImageOrVideo(&processedFiles)
	if !isImageOrVideoCreated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create your media files"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "media files have been successfully posted"})
}

func GetPropertyImageOrVideoById(c *gin.Context) {
	propertyImageOrVideo := imagesOrVideosRepo.GetPropertyImageOrVideoById(c.Param("id"))
	if propertyImageOrVideo == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this media file does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyImageOrVideoResponse(*propertyImageOrVideo)})
}

func UpdatePropertyImageOrVideo(c *gin.Context) {
	var imageOrVideoUpdateDetails imagesOrVideosDtos.PropertyImageOrVideoUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&imageOrVideoUpdateDetails)

	modelFieldsValidationError := validateModelFields.Struct(imageOrVideoUpdateDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	fileUrl := <-storage.UploadFile(imageOrVideoUpdateDetails.File, imageOrVideoUpdateDetails.Name, c)

	imageOrVideoUpdate := propertyModels.PropertyImageOrVideo{
		Id:           imageOrVideoUpdateDetails.Id,
		PropertyId:   imageOrVideoUpdateDetails.PropertyId,
		Size:         imageOrVideoUpdateDetails.Size,
		FileType:     imageOrVideoUpdateDetails.FileType,
		ContentType:  imageOrVideoUpdateDetails.ContentType,
		Uri:          fileUrl,
		Name:         imageOrVideoUpdateDetails.Name,
		PropertyType: imageOrVideoUpdateDetails.PropertyType,
	}

	isImageOrVideoUpdated := imagesOrVideosRepo.UpdatePropertyImageOrVideo(&imageOrVideoUpdate)
	if !isImageOrVideoUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your media file"})
		return
	}
	generalServices.GetPropertyTypeById(c, imageOrVideoUpdateDetails.PropertyType, imageOrVideoUpdate.PropertyId)
}

func DeletePropertyImageOrVideo(c *gin.Context) {
	imageOrVideo := imagesOrVideosRepo.GetPropertyImageOrVideoById(c.Param("id"))
	if imageOrVideo == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this media file does not exist"})
		return
	}
	<-storage.DeleteFile(imageOrVideo.Name, c)
	isImageOrVideoDeleted := imagesOrVideosRepo.DeletePropertyImageOrVideo(c.Param("id"))
	if !isImageOrVideoDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this media file does not exist"})
		return
	} else {
		generalServices.GetPropertyTypeById(c, imageOrVideo.PropertyType, imageOrVideo.PropertyId)
	}
}
