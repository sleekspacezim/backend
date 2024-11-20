package admin

import (
	"net/http"

	commercialDtos "SleekSpace/dtos/property/commercial"
	landDtos "SleekSpace/dtos/property/land"
	propertyLocationDtos "SleekSpace/dtos/property/location"
	residentialDtos "SleekSpace/dtos/property/residential"
	standDtos "SleekSpace/dtos/property/stand"
	commercialRepo "SleekSpace/repositories/property/commercial"
	propertyInsightsRepo "SleekSpace/repositories/property/insights"
	landRepo "SleekSpace/repositories/property/land"
	propertyLocationRepo "SleekSpace/repositories/property/location"
	propertyMediaRepo "SleekSpace/repositories/property/media"
	residentialRepo "SleekSpace/repositories/property/residential"
	standRepo "SleekSpace/repositories/property/stand"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
)

func GetAllResidentialRentalProperties(c *gin.Context) {
	residentialRentalProperties := residentialRepo.GetAllResidentialRentalProperties(c)
	responseList := []residentialDtos.ResidentialPropertyForRentResponseDto{}
	if len(residentialRentalProperties) > 0 {
		for i := 0; i < len(residentialRentalProperties); i++ {
			responseItem := propertyUtilities.ResidentialRentalPropertyResponse(residentialRentalProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": responseList,
			"totalPages": c.GetInt("totalPages"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
	})
}

func GetAllStands(c *gin.Context) {
	stands := standRepo.GetAllStands(c)
	responseList := []standDtos.StandResponseDTO{}
	if len(stands) > 0 {
		for i := 0; i < len(stands); i++ {
			responseItem := propertyUtilities.PropertyStandResponse(stands[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": responseList,
			"totalPages": c.GetInt("totalPages"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
	})
}

func GetLandWithManager(c *gin.Context) {
	result := landRepo.GetLandPropertyForSaleByIdWithmanager(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})
}

func GetAllLandProperties(c *gin.Context) {
	landProperties := landRepo.GetAllLandPropertiesForSale(c)
	responseList := []landDtos.LandForSalePropertyWithManagerResponseDto{}
	if len(landProperties) > 0 {
		for i := 0; i < len(landProperties); i++ {
			responseItem := propertyUtilities.LandPropertyWithManagerResponse(landProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": responseList,
			"totalPages": c.GetInt("totalPages"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
	})
}

func GetAllResidentialForSaleProperties(c *gin.Context) {
	residentialPropertiesForSale := residentialRepo.GetAllResidentialPropertiesForSale(c)
	responseList := []residentialDtos.ResidentialPropertyForSaleResponseDto{}
	if len(residentialPropertiesForSale) > 0 {
		for i := 0; i < len(residentialPropertiesForSale); i++ {
			responseItem := propertyUtilities.ResidentialForSalePropertyResponse(residentialPropertiesForSale[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": responseList,
			"totalPages": c.GetInt("totalPages"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
	})
}

func GetAllCommercialRentalProperties(c *gin.Context) {
	commercialRentalProperties := commercialRepo.GetAllCommercialRentalProperties(c)
	responseList := []commercialDtos.CommercialForRentPropertyResponseDto{}
	if len(commercialRentalProperties) > 0 {
		for i := 0; i < len(commercialRentalProperties); i++ {
			responseItem := propertyUtilities.CommercialPropertyForRentResponse(commercialRentalProperties[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": responseList,
			"totalPages": c.GetInt("totalPages"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
	})
}

func GetAllCommercialForSaleProperties(c *gin.Context) {
	commercialPropertiesForSale := commercialRepo.GetAllCommercialPropertiesForSale(c)
	responseList := []commercialDtos.CommercialForSalePropertyResponseDto{}
	if len(commercialPropertiesForSale) > 0 {
		for i := 0; i < len(commercialPropertiesForSale); i++ {
			responseItem := propertyUtilities.CommercialPropertyForSaleResponse(commercialPropertiesForSale[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"properties": responseList,
			"totalPages": c.GetInt("totalPages"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"properties": responseList,
		"totalPages": c.GetInt("totalPages"),
	})
}

func GetAllPropertiesImagesOrVideos(c *gin.Context) {
	imagesOrVideosList := propertyMediaRepo.GetAllPropertiesImagesOrVideos()
	c.JSON(http.StatusOK, gin.H{
		"response": imagesOrVideosList,
	})
}

func GetAllPropertiesInsights(c *gin.Context) {
	insightsList := propertyInsightsRepo.GetAllPropertiesInsights()
	c.JSON(http.StatusOK, gin.H{
		"response": insightsList,
	})
}

func GetAllPropertiesLocation(c *gin.Context) {
	propertyLocations := propertyLocationRepo.GetAllPropertyLocations()
	responseList := []propertyLocationDtos.PropertyLocationUpdateAndResponseDto{}
	if len(propertyLocations) > 0 {
		for i := 0; i < len(propertyLocations); i++ {
			responseItem := propertyUtilities.PropertyLocationResponse(propertyLocations[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"response": responseList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": responseList,
	})
}
