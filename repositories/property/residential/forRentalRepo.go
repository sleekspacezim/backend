package residential

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	pagination "SleekSpace/repositories"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateResidentialRentalProperty(property *managerModels.ResidentialRentalProperty) bool {
	err := db.DB.Create(property)
	if err != nil {
		println(err)
	}
	return true
}

func GetResidentialRentalPropertyById(id string) *managerModels.ResidentialRentalProperty {
	var property managerModels.ResidentialRentalProperty
	result := db.DB.First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetResidentialRentalPropertyWithAllAssociationsById(id string) *managerModels.ResidentialRentalProperty {
	var property managerModels.ResidentialRentalProperty
	result := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetResidentialRentalPropertyWithAllAssociationsByUniqueId(uniqueId string) *managerModels.ResidentialRentalProperty {
	var property managerModels.ResidentialRentalProperty
	result := db.DB.Where("unique_id= ?", uniqueId).
		Preload(clause.Associations).
		First(&property)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetManagerResidentialRentalPropertiesByManagerId(managerId string) []managerModels.ResidentialRentalProperty {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("ResidentialRentalProperty").
		First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.ResidentialRentalProperty
}

func GetAllResidentialRentalProperties(c *gin.Context) []managerModels.ResidentialRentalProperty {
	var properties = []managerModels.ResidentialRentalProperty{}
	err := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Order("created_at DESC, id DESC").
		Scopes(pagination.Paginate(c)).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateResidentialRentalProperty(update *managerModels.ResidentialRentalProperty) bool {
	db.DB.Save(update)
	return true
}

func DeleteResidentialRentalPropertyById(id string) bool {
	property := GetResidentialRentalPropertyById(id)
	if property == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&property)
	return true
}
