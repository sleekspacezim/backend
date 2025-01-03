package commercial

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

func CreateCommercialRentalProperty(property *managerModels.CommercialRentalProperty) bool {
	err := db.DB.Create(property)
	if err != nil {
		println(err)
	}
	return true
}

func GetCommercialRentalPropertyById(id string) *managerModels.CommercialRentalProperty {
	var property managerModels.CommercialRentalProperty
	result := db.DB.First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetCommercialRentalPropertyWithAllAssociationsById(id string) *managerModels.CommercialRentalProperty {
	var property managerModels.CommercialRentalProperty
	result := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetCommercialRentalPropertyWithAllAssociationsByUniqueId(uniqueId string) *managerModels.CommercialRentalProperty {
	var property managerModels.CommercialRentalProperty
	result := db.DB.Where("unique_id= ?", uniqueId).Preload(clause.Associations).First(&property)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetManagerCommercialRentalPropertiesByManagerId(managerId string) []managerModels.CommercialRentalProperty {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("CommercialRentalProperty").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.CommercialRentalProperty
}

func GetAllCommercialRentalProperties(c *gin.Context) []managerModels.CommercialRentalProperty {
	var properties = []managerModels.CommercialRentalProperty{}
	err := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Scopes(
			scopes.SortProperties(c),
			scopes.RentFilter(c),
			scopes.PropertyStructureTypeFilter(c),
			scopes.CurrencyFilter(c),
			scopes.PropertySizeFilter(c),
			scopes.NumberOfRoomsFilter(c),
			scopes.NumberOfRoomsToLetFilter(c),
			pagination.Paginate(c),
		).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func GetAllCommercialRentalPropertiesByLocation(c *gin.Context, location string) []managerModels.CommercialRentalProperty {
	var properties = []managerModels.CommercialRentalProperty{}
	err := db.DB.
		Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Joins("JOIN property_locations ON property_locations.property_id = commercial_rental_properties.unique_id").
		Where("property_locations.display_name ILIKE ?", "%"+location+"%").
		Scopes(pagination.Paginate(c), scopes.SortProperties(c)).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateCommercialRentalProperty(update *managerModels.CommercialRentalProperty) bool {
	db.DB.Save(update)
	return true
}

func DeleteCommercialRentalPropertyById(id string) bool {
	property := GetCommercialRentalPropertyById(id)
	if property == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&property)
	return true
}
