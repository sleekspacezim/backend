package property

import (
	"net/http"

	commercialRepo "SleekSpace/repositories/property/commercial"
	landRepo "SleekSpace/repositories/property/land"
	residentialRepo "SleekSpace/repositories/property/residential"
	standRepo "SleekSpace/repositories/property/stand"
	constants "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
)

func GetPropertyTypeById(c *gin.Context, propertyType string, propertyId int) {
	if propertyType == constants.StandPropertyType {
		stand := standRepo.GetStandWithAllAssociationsByUniqueId(generalUtilities.ConvertIntToString(propertyId))
		if stand == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyStandResponse(*stand)})
		return
	} else if propertyType == constants.LandPropertyType {
		land := landRepo.GetLandPropertyForSaleWithAllAssociationsByUniqueId(generalUtilities.ConvertIntToString(propertyId))
		if land == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.LandPropertyResponse(*land)})
		return
	} else if propertyType == constants.ResidentialRentalPropertyType {
		property := residentialRepo.GetResidentialRentalPropertyWithAllAssociationsByUniqueId(generalUtilities.ConvertIntToString(propertyId))
		if property == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ResidentialRentalPropertyResponse(*property)})
		return
	} else if propertyType == constants.ResidentialPropertyForSaleType {
		property := residentialRepo.GetResidentialPropertyForSaleWithAllAssociationsByUniqueId(generalUtilities.ConvertIntToString(propertyId))
		if property == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ResidentialForSalePropertyResponse(*property)})
		return
	} else if propertyType == constants.CommercialRentalPropertyType {
		property := commercialRepo.GetCommercialRentalPropertyWithAllAssociationsByUniqueId(generalUtilities.ConvertIntToString(propertyId))
		if property == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForRentResponse(*property)})
		return
	} else {
		property := commercialRepo.GetCommercialPropertyForSaleWithAllAssociationsByUniqueId(generalUtilities.ConvertIntToString(propertyId))
		if property == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForSaleResponse(*property)})
		return
	}
}
