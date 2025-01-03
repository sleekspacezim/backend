package land

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	pagination "SleekSpace/repositories"
	scopes "SleekSpace/repositories/scopes"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateLandPropertyForSale(land *managerModels.LandForSaleProperty) bool {
	err := db.DB.Create(land)
	if err != nil {
		println(err)
	}
	return true
}

func GetLandPropertyForSaleById(id string) *managerModels.LandForSaleProperty {
	var land managerModels.LandForSaleProperty
	result := db.DB.First(&land, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &land
}

func GetLandPropertyForSaleByIdWithmanager(id string) *managerModels.Manager {
	var land managerModels.Manager
	result := db.DB.
		Joins("JOIN manager_profile_pictures ON manager_profile_pictures.manager_id = managers.id").
		Where("managers.id =?", id).
		Find(&land)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &land
}

func GetLandPropertyForSaleWithAllAssociationsByUniqueId(uniqueId string) *managerModels.LandForSaleProperty {
	var land managerModels.LandForSaleProperty
	result := db.DB.Where("unique_id= ?", uniqueId).Preload(clause.Associations).First(&land)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &land
}

func GetLandPropertyForSaleWithAllAssociationsById(id string) *managerModels.LandForSaleProperty {
	var land managerModels.LandForSaleProperty
	result := db.DB.Preload(clause.Associations).Preload("Manager.ProfilePicture").Preload("Manager.ManagerContactNumbers").First(&land, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &land
}

func GetManagerLandPropertiesForSaleByManagerId(managerId string) []managerModels.LandForSaleProperty {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("LandForSaleProperty").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.LandForSaleProperty
}

func GetAllLandPropertiesForSale(c *gin.Context) []managerModels.LandForSaleProperty {
	var properties = []managerModels.LandForSaleProperty{}
	err := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Scopes(
			scopes.SortProperties(c),
			scopes.PriceFilter(c),
			scopes.PropertyStructureTypeFilter(c),
			scopes.CurrencyFilter(c),
			scopes.PropertySizeFilter(c),
			pagination.Paginate(c),
		).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func GetAllLandPropertiesByLocation(
	c *gin.Context,
	location string,
) []managerModels.LandForSaleProperty {
	var properties = []managerModels.LandForSaleProperty{}
	err := db.DB.
		Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Joins("JOIN property_locations ON property_locations.property_id = land_for_sale_properties.unique_id").
		Where("property_locations.display_name ILIKE ?", "%"+location+"%").
		Scopes(pagination.Paginate(c), scopes.SortProperties(c)).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateLandPropertyForSale(update *managerModels.LandForSaleProperty) bool {
	db.DB.Save(update)
	return true
}

func DeleteLandPropertyForSaleById(id string) bool {
	land := GetLandPropertyForSaleById(id)
	if land == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&land)
	return true
}
