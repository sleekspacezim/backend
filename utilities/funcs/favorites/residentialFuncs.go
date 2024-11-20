package favorites

import (
	residentialDtos "SleekSpace/dtos/property/residential"
	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"

	"github.com/gin-gonic/gin"
)

func ProcessFavoritesForResidentialRentalPropertyWithManager(
	responseList []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto, c *gin.Context,
) []residentialDtos.ResidentialPropertyForRentWithManagerResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteResidentialRentalProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteResidentialRentalProperties); i++ {
			idsMap[user.FavoriteResidentialRentalProperties[i]] = true
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

func ProcessFavoritesForResidentialRentalPropertyWithoutManager(
	responseList []residentialDtos.ResidentialPropertyForRentResponseDto, c *gin.Context,
) []residentialDtos.ResidentialPropertyForRentResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteResidentialRentalProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteResidentialRentalProperties); i++ {
			idsMap[user.FavoriteResidentialRentalProperties[i]] = true
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

func ProcessFavoritesForResidentialForSalePropertyWithManager(
	responseList []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto, c *gin.Context,
) []residentialDtos.ResidentialPropertyForSaleWithManagerResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteResidentialForSaleProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteResidentialForSaleProperties); i++ {
			idsMap[user.FavoriteResidentialForSaleProperties[i]] = true
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

func ProcessFavoritesForResidentialForSalePropertyWithoutManager(
	responseList []residentialDtos.ResidentialPropertyForSaleResponseDto, c *gin.Context,
) []residentialDtos.ResidentialPropertyForSaleResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteResidentialForSaleProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteResidentialForSaleProperties); i++ {
			idsMap[user.FavoriteResidentialForSaleProperties[i]] = true
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
