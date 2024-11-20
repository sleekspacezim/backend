package dtos

type UserResponseDTO struct {
	Id                                   int                           `json:"id"`
	FavoriteLandProperties               []int                         `json:"favoriteLandProperties"`
	FavoriteStands                       []int                         `json:"favoriteStands" gorm:"serializer:json"`
	FavoriteCommercialForSaleProperties  []int                         `json:"favoriteCommercialForSaleProperties"`
	FavoriteCommercialRentalProperties   []int                         `json:"favoriteCommercialRentalProperties"`
	FavoriteResidentialForSaleProperties []int                         `json:"favoriteResidentialForSaleProperties"`
	FavoriteResidentialRentalProperties  []int                         `json:"favoriteResidentialRentalProperties"`
	Email                                string                        `json:"email"`
	GivenName                            string                        `json:"givenName"`
	FamilyName                           string                        `json:"familyName"`
	Role                                 string                        `json:"role"`
	AccessToken                          string                        `json:"accessToken"`
	ProfilePicture                       UserProfilePictureResponseDTO `json:"profilePicture"`
	Location                             LocationDTO                   `json:"location"`
	ContactNumbers                       []ContactNumberDTO            `json:"contactNumbers"`
}
