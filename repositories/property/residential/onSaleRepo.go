package residential

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	pagination "SleekSpace/repositories"
	sort "SleekSpace/repositories/scopes"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateResidentialPropertyForSale(property *managerModels.ResidentialPropertyForSale) bool {
	err := db.DB.Create(property)
	if err != nil {
		println(err)
	}
	return true
}

func GetResidentialPropertyForSaleById(id string) *managerModels.ResidentialPropertyForSale {
	var property managerModels.ResidentialPropertyForSale
	result := db.DB.First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetResidentialPropertyForSaleWithAllAssociationsById(id string) *managerModels.ResidentialPropertyForSale {
	var property managerModels.ResidentialPropertyForSale
	result := db.DB.Preload(clause.Associations).Preload("Manager.ProfilePicture").Preload("Manager.ManagerContactNumbers").First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetResidentialPropertyForSaleWithAllAssociationsByUniqueId(uniqueId string) *managerModels.ResidentialPropertyForSale {
	var property managerModels.ResidentialPropertyForSale
	result := db.DB.Where("unique_id= ?", uniqueId).Preload(clause.Associations).First(&property)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetManagerResidentialPropertiesForSaleByManagerId(managerId string) []managerModels.ResidentialPropertyForSale {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("ResidentialPropertyForSale").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.ResidentialPropertyForSale
}

func GetAllResidentialPropertiesForSale(c *gin.Context) []managerModels.ResidentialPropertyForSale {
	var properties = []managerModels.ResidentialPropertyForSale{}
	err := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Scopes(pagination.Paginate(c), sort.SortProperties(c)).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateResidentialPropertyForSale(update *managerModels.ResidentialPropertyForSale) bool {
	db.DB.Save(update)
	return true
}

func DeleteResidentialPropertyForSaleById(id string) bool {
	property := GetResidentialPropertyForSaleById(id)
	if property == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&property)
	return true
}