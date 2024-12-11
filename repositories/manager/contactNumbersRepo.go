package repositories

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	"errors"

	"gorm.io/gorm"
)

func GetManagerContactNumbersByManagerId(managerId int) []managerModels.ManagerContactNumber {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("ManagerContactNumbers").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.ManagerContactNumbers
}

func UpdateManagerContactNumbers(
	manager *managerModels.Manager,
	updateManagerContactNumbersList []managerModels.ManagerContactNumber,
) bool {
	manager.ManagerContactNumbers = updateManagerContactNumbersList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&manager)
	return true
}
