package models

import (
	baseModel "SleekSpace/models"
	managerModels "SleekSpace/models/manager"
)

type User struct {
	baseModel.MyModel
	Id                                   int                   `json:"id" gorm:"primary_key"`
	FavoriteLandProperties               []int                 `json:"favoriteLandProperties" gorm:"serializer:json"`
	FavoriteStands                       []int                 `json:"favoriteStands" gorm:"serializer:json"`
	FavoriteCommercialForSaleProperties  []int                 `json:"favoriteCommercialForSaleProperties" gorm:"serializer:json"`
	FavoriteCommercialRentalProperties   []int                 `json:"favoriteCommercialRentalProperties" gorm:"serializer:json"`
	FavoriteResidentialForSaleProperties []int                 `json:"favoriteResidentialForSaleProperties" gorm:"serializer:json"`
	FavoriteResidentialRentalProperties  []int                 `json:"favoriteResidentialRentalProperties" gorm:"serializer:json"`
	IsActive                             bool                  `json:"isActive"`
	IsSocialsAuthenticated               bool                  `json:"isSocialsAuthenticated"`
	GivenName                            string                `json:"givenName" validate:"required,min=2,max=50"`
	FamilyName                           string                `json:"familyName" validate:"required,min=2,max=50"`
	Email                                string                `json:"email" gorm:"unique" validate:"email,required"`
	Password                             string                `json:"password"`
	AccessToken                          string                `json:"accessToken"`
	SocialMediaProvider                  string                `json:"socialMediaProvider"`
	Role                                 string                `json:"role"`
	Location                             Location              `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RegistrationCode                     VerificationCode      `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Manager                              managerModels.Manager `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	ProfilePicture UserProfilePicture `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ContactNumbers []ContactNumber    `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
