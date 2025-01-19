package stand

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

func CreateStandForSale(stand *managerModels.Stand) bool {
	err := db.DB.Create(stand)
	if err != nil {
		println(err)
	}
	return true
}

func GetStandById(id string) *managerModels.Stand {
	var stand managerModels.Stand
	result := db.DB.First(&stand, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &stand
}

func GetStandWithAllAssociationsById(id string) *managerModels.Stand {
	var stand managerModels.Stand
	result := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		First(&stand, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &stand
}

func GetStandWithAllAssociationsByUniqueId(uniqueId string) *managerModels.Stand {
	var stand managerModels.Stand
	result := db.DB.Where("unique_id= ?", uniqueId).Preload(clause.Associations).First(&stand)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &stand
}

func GetManagerStandsByManagerId(managerId string) []managerModels.Stand {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("Stand").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.Stand
}

func GetStandPropertyForSaleByIds(ids []int, c *gin.Context) []managerModels.Stand {
	var properties = []managerModels.Stand{}
	result := db.DB.Where("id IN ?", ids).
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

func GetAllStands(c *gin.Context) []managerModels.Stand {
	var stands = []managerModels.Stand{}
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
		Find(&stands)
	if err != nil {
		println(err.Error, err.Name())
	}
	return stands
}

func GetAllStandsByLocation(
	c *gin.Context,
	location string,
) []managerModels.Stand {
	var properties = []managerModels.Stand{}
	err := db.DB.
		Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Joins("JOIN property_locations ON property_locations.property_id = stands.unique_id").
		Where("property_locations.display_name ILIKE ?", "%"+location+"%").
		Scopes(
			scopes.SortProperties(c),
			pagination.Paginate(c),
		).
		Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateStand(update *managerModels.Stand) bool {
	db.DB.Save(update)
	return true
}

func DeleteStandById(id string) bool {
	stand := GetStandById(id)
	if stand == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&stand)
	return true
}
