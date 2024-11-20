package favorites

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	pagination "SleekSpace/repositories"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteResidentialRentalProperties(
	residentialRentalPropertiesIds []int,
	c *gin.Context,
) []managerModels.ResidentialRentalProperty {
	var properties = []managerModels.ResidentialRentalProperty{}
	result := db.DB.Where(
		"id IN ?", residentialRentalPropertiesIds).
		Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Order("created_at DESC, id DESC").
		Scopes(pagination.Paginate(c)).
		Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
