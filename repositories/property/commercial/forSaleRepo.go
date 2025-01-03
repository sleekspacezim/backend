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

func CreateCommercialPropertyForSale(property *managerModels.CommercialForSaleProperty) bool {
	err := db.DB.Create(property)
	if err != nil {
		println(err)
	}
	return true
}

func GetCommercialPropertyForSaleById(id string) *managerModels.CommercialForSaleProperty {
	var property managerModels.CommercialForSaleProperty
	result := db.DB.First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetCommercialPropertyForSaleWithAllAssociationsByUniqueId(uniqueId string) *managerModels.CommercialForSaleProperty {
	var property managerModels.CommercialForSaleProperty
	result := db.DB.Where("unique_id= ?", uniqueId).
		Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		First(&property)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetCommercialPropertyForSaleWithAllAssociationsById(id string) *managerModels.CommercialForSaleProperty {
	var property managerModels.CommercialForSaleProperty
	result := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetManagerCommercialPropertiesForSaleByManagerId(managerId string) []managerModels.CommercialForSaleProperty {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("CommercialForSaleProperty").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.CommercialForSaleProperty
}

func GetAllCommercialPropertiesForSale(c *gin.Context) []managerModels.CommercialForSaleProperty {
	var properties = []managerModels.CommercialForSaleProperty{}
	err := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Scopes(
			scopes.SortProperties(c),
			scopes.PriceFilter(c),
			scopes.PropertyStructureTypeFilter(c),
			scopes.CurrencyFilter(c),
			scopes.PropertySizeFilter(c),
			scopes.NumberOfRoomsFilter(c),
			pagination.Paginate(c),
		).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func GetAllCommercialPropertiesForSaleByLocation(
	c *gin.Context,
	location string,
) []managerModels.CommercialForSaleProperty {
	var properties = []managerModels.CommercialForSaleProperty{}
	err := db.DB.
		Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Joins("JOIN property_locations ON property_locations.property_id = commercial_for_sale_properties.unique_id").
		Where("property_locations.display_name ILIKE ?", "%"+location+"%").
		Scopes(pagination.Paginate(c), scopes.SortProperties(c)).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateCommercialPropertyForSale(update *managerModels.CommercialForSaleProperty) bool {
	db.DB.Save(update)
	return true
}

func DeleteCommercialPropertyForSaleById(id string) bool {
	property := GetCommercialPropertyForSaleById(id)
	if property == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&property)
	return true
}
