package stand

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	pagination "SleekSpace/repositories"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateStandForSale(stand *managerModels.PropertyStand) bool {
	err := db.DB.Create(stand)
	if err != nil {
		println(err)
	}
	return true
}

func GetStandById(id string) *managerModels.PropertyStand {
	var stand managerModels.PropertyStand
	result := db.DB.First(&stand, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &stand
}

func GetStandWithAllAssociationsById(id string) *managerModels.PropertyStand {
	var stand managerModels.PropertyStand
	result := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		First(&stand, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &stand
}

func GetStandWithAllAssociationsByUniqueId(uniqueId string) *managerModels.PropertyStand {
	var stand managerModels.PropertyStand
	result := db.DB.Where("unique_id= ?", uniqueId).Preload(clause.Associations).First(&stand)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &stand
}

func GetManagerStandsByManagerId(managerId string) []managerModels.PropertyStand {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("PropertyStand").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.PropertyStand
}

func GetAllStands(c *gin.Context) []managerModels.PropertyStand {
	var stands = []managerModels.PropertyStand{}
	err := db.DB.Preload(clause.Associations).
		Preload("Manager.ProfilePicture").
		Preload("Manager.ManagerContactNumbers").
		Order("created_at DESC, id DESC").
		Scopes(pagination.Paginate(c)).Find(&stands)
	if err != nil {
		println(err.Error, err.Name())
	}
	return stands
}

func UpdateStand(update *managerModels.PropertyStand) bool {
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
