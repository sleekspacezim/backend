package favorites

import (
	commercialDtos "SleekSpace/dtos/property/commercial"
	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"

	"github.com/gin-gonic/gin"
)

func ProcessFavoritesForCommercialRentalPropertyWithManager(
	responseList []commercialDtos.CommercialForRentPropertyWithManagerResponseDto, c *gin.Context,
) []commercialDtos.CommercialForRentPropertyWithManagerResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteCommercialRentalProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteCommercialRentalProperties); i++ {
			idsMap[user.FavoriteCommercialRentalProperties[i]] = true
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

func ProcessFavoritesForCommercialRentalPropertyWithoutManager(
	responseList []commercialDtos.CommercialForRentPropertyResponseDto, c *gin.Context,
) []commercialDtos.CommercialForRentPropertyResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteCommercialRentalProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteCommercialRentalProperties); i++ {
			idsMap[user.FavoriteCommercialRentalProperties[i]] = true
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

func ProcessFavoritesForCommercialForSalePropertyWithManager(
	responseList []commercialDtos.CommercialForSalePropertyWithManagerResponseDto, c *gin.Context,
) []commercialDtos.CommercialForSalePropertyWithManagerResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteCommercialForSaleProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteCommercialForSaleProperties); i++ {
			idsMap[user.FavoriteCommercialForSaleProperties[i]] = true
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

func ProcessFavoritesForCommercialForSalePropertyWithoutManager(
	responseList []commercialDtos.CommercialForSalePropertyResponseDto, c *gin.Context,
) []commercialDtos.CommercialForSalePropertyResponseDto {
	email := c.MustGet("user").(*userModels.User).Email
	user := userRepo.GetUserByEmail(email)
	idsMap := make(map[int]bool)
	if len(user.FavoriteCommercialForSaleProperties) > 0 && len(responseList) > 0 {
		for i := 0; i < len(user.FavoriteCommercialForSaleProperties); i++ {
			idsMap[user.FavoriteCommercialForSaleProperties[i]] = true
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
