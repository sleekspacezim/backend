package favorites

import (
	landDtos "SleekSpace/dtos/property/land"
	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"

	"github.com/gin-gonic/gin"
)

func ProcessFavoritesForLandPropertyWithManager(
	responseList []landDtos.LandForSalePropertyWithManagerResponseDto, c *gin.Context,
) []landDtos.LandForSalePropertyWithManagerResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteLandProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteLandProperties); i++ {
			idsMap[user.FavoriteLandProperties[i]] = true
		}
		for i := 0; i < len(responseList); i++ {
			if !idsMap[responseList[i].Id] {
				responseList[i].IsFavorite = false
			} else {
				responseList[i].IsFavorite = true
			}
		}
	}
	return responseList
}

func ProcessFavoritesForLandPropertyWithoutManager(
	responseList []landDtos.LandForSalePropertyResponseDto, c *gin.Context,
) []landDtos.LandForSalePropertyResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteLandProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteLandProperties); i++ {
			idsMap[user.FavoriteLandProperties[i]] = true
		}
		for i := 0; i < len(responseList); i++ {
			if !idsMap[responseList[i].Id] {
				responseList[i].IsFavorite = false
			} else {
				responseList[i].IsFavorite = true
			}
		}
	}
	return responseList
}
