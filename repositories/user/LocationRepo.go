package repositories

import (
	"SleekSpace/db"
	userModels "SleekSpace/models/user"
)

func CreateLocation(user *userModels.User, location *userModels.Location) bool {
	err := db.DB.Model(user).Association("Location").Append(location)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetAllUsersLocations() []userModels.Location {
	var locations = []userModels.Location{}
	err := db.DB.Find(&locations)
	if err != nil {
		println(err.Error, err.Name())
	}
	return locations
}

func GetLocationByUserId(userId string) userModels.Location {
	var user userModels.User
	err := db.DB.Preload("Location").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Location
}

func GetLocationById(locationId int) userModels.Location {
	var location = userModels.Location{}
	db.DB.First(&location, locationId)
	return location
}

func UpdateLocation(locationUpdate *userModels.Location) bool {
	db.DB.Save(locationUpdate)
	return true
}

func DeleteLocation(locationId int) bool {
	db.DB.Unscoped().Delete(&userModels.Location{}, locationId)
	return true
}
