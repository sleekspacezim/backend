package db

import (
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	userModels "SleekSpace/models/user"
	generalUtilities "SleekSpace/utilities/funcs/general"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	databaseDetails := generalUtilities.GetEnvVariables().DatabaseDetails
	db, err := gorm.Open(postgres.Open(databaseDetails), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&propertyModels.PropertyInsights{},
		&propertyModels.PropertyImageOrVideo{},
		&propertyModels.PropertyLocation{},
		&propertyModels.PropertyReport{},
	)

	db.AutoMigrate(
		&managerModels.CommercialForSaleProperty{},
		&managerModels.CommercialRentalProperty{},
		&managerModels.LandForSaleProperty{},
		&managerModels.ResidentialPropertyForSale{},
		&managerModels.ResidentialRentalProperty{},
		&managerModels.PropertyStand{},
	)

	db.AutoMigrate(
		&managerModels.ManagerContactNumber{},
		&managerModels.ManagerProfilePicture{},
	)

	db.AutoMigrate(
		&userModels.VerificationCode{},
		&userModels.ContactNumber{},
		&userModels.Location{},
		&managerModels.Manager{},
		&userModels.UserProfilePicture{},
	)

	db.AutoMigrate(&userModels.User{})

	DB = db
}
