package favorites

import (
	standDtos "SleekSpace/dtos/property/stand"
	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"

	"github.com/gin-gonic/gin"
)

func ProcessFavoritesForStandPropertyWithManager(
	responseList []standDtos.StandWithManagerResponseDTO, c *gin.Context,
) []standDtos.StandWithManagerResponseDTO {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteStands) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteStands); i++ {
			idsMap[user.FavoriteStands[i]] = true
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

func ProcessFavoritesForStandPropertyWithoutManager(
	responseList []standDtos.StandResponseDTO, c *gin.Context,
) []standDtos.StandResponseDTO {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteStands) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteStands); i++ {
			idsMap[user.FavoriteStands[i]] = true
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
