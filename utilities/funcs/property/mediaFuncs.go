package property

import (
	propertyMediaDtos "SleekSpace/dtos/property/media"
	propertyModels "SleekSpace/models/property"
)

type MediaFile struct {
	FileSrc string
	Name    string
}

func MediaListWithNoPropertyId(mediaList []propertyMediaDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto) []MediaFile {
	mediaFiles := []MediaFile{}
	if len(mediaList) == 0 {
		return mediaFiles
	}
	for i := 0; i < len(mediaList); i++ {
		file := MediaFile{
			FileSrc: mediaList[i].File,
			Name:    mediaList[i].Name,
		}
		mediaFiles = append(mediaFiles, file)
	}
	return mediaFiles
}

func MediaListWithPropertyId(mediaList []propertyMediaDtos.PropertyImageOrVideoCreationWithPropertyIdDto) []MediaFile {
	mediaFiles := []MediaFile{}
	if len(mediaList) == 0 {
		return mediaFiles
	}
	for i := 0; i < len(mediaList); i++ {
		file := MediaFile{
			FileSrc: mediaList[i].File,
			Name:    mediaList[i].Name,
		}
		mediaFiles = append(mediaFiles, file)
	}
	return mediaFiles
}

func ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(
	propertyMediaList []propertyMediaDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto,
	propertyType string,
	mediaUrls map[string]string,
) []propertyModels.PropertyImageOrVideo {
	mediaList := []propertyModels.PropertyImageOrVideo{}
	if len(propertyMediaList) > 0 {
		for i := 0; i < len(propertyMediaList); i++ {
			media := propertyModels.PropertyImageOrVideo{
				Uri:          mediaUrls[propertyMediaList[i].Name],
				FileType:     propertyMediaList[i].FileType,
				ContentType:  propertyMediaList[i].ContentType,
				Size:         propertyMediaList[i].Size,
				Name:         propertyMediaList[i].Name,
				PropertyType: propertyType,
			}
			mediaList = append(mediaList, media)
		}
	}
	return mediaList
}

func ConvertPropertyImagesOrVideosWithPropertyIdToModel(propertyMediaList []propertyMediaDtos.PropertyImageOrVideoCreationWithPropertyIdDto, mediaUrls map[string]string) []propertyModels.PropertyImageOrVideo {
	mediaList := []propertyModels.PropertyImageOrVideo{}
	if len(propertyMediaList) > 0 {
		for i := 0; i < len(propertyMediaList); i++ {
			media := propertyModels.PropertyImageOrVideo{
				Uri:          mediaUrls[propertyMediaList[i].Name],
				PropertyId:   propertyMediaList[i].PropertyId,
				FileType:     propertyMediaList[i].FileType,
				ContentType:  propertyMediaList[i].ContentType,
				Size:         propertyMediaList[i].Size,
				Name:         propertyMediaList[i].Name,
				PropertyType: propertyMediaList[i].PropertyType,
			}
			mediaList = append(mediaList, media)
		}
	}
	return mediaList
}

func ProcessedPropertyImageAndVideosListToResponse(mediaList []propertyModels.PropertyImageOrVideo) []propertyMediaDtos.PropertyImageOrVideoResponseDto {
	dtoList := []propertyMediaDtos.PropertyImageOrVideoResponseDto{}
	if len(mediaList) > 0 {
		for i := 0; i < len(mediaList); i++ {
			dto := propertyMediaDtos.PropertyImageOrVideoResponseDto{
				Id:           mediaList[i].Id,
				PropertyId:   mediaList[i].PropertyId,
				Uri:          mediaList[i].Uri,
				FileType:     mediaList[i].FileType,
				ContentType:  mediaList[i].ContentType,
				Size:         mediaList[i].Size,
				Name:         mediaList[i].Name,
				PropertyType: mediaList[i].PropertyType,
			}
			dtoList = append(dtoList, dto)
		}
	}
	return dtoList
}

func PropertyImageOrVideoResponse(media propertyModels.PropertyImageOrVideo) propertyMediaDtos.PropertyImageOrVideoResponseDto {
	return propertyMediaDtos.PropertyImageOrVideoResponseDto{
		Id:           media.Id,
		PropertyId:   media.PropertyId,
		Uri:          media.Uri,
		FileType:     media.FileType,
		ContentType:  media.ContentType,
		Size:         media.Size,
		Name:         media.Name,
		PropertyType: media.PropertyType,
	}
}
