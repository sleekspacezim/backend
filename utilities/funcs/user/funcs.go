package utilities

import (
	userDtos "SleekSpace/dtos/user"
	userModels "SleekSpace/models/user"
)

func processedContactNumbers(contactNumbers []userModels.ContactNumber) []userDtos.ContactNumberDTO {
	contacts := []userDtos.ContactNumberDTO{}
	for i := 0; i < len(contactNumbers); i++ {
		contact := userDtos.ContactNumberDTO{
			Id:           contactNumbers[i].Id,
			CountryAbbrv: contactNumbers[i].CountryAbbrv,
			CountryCode:  contactNumbers[i].CountryCode,
			Number:       contactNumbers[i].Number,
			Type:         contactNumbers[i].Type,
			UserId:       contactNumbers[i].UserId,
		}
		contacts = append(contacts, contact)
	}
	return contacts
}

func UserResponseMapper(user *userModels.User, accessToken string) userDtos.UserResponseDTO {
	return userDtos.UserResponseDTO{
		Id:                                   user.Id,
		Email:                                user.Email,
		ContactNumbers:                       processedContactNumbers(user.ContactNumbers),
		FamilyName:                           user.FamilyName,
		GivenName:                            user.GivenName,
		AccessToken:                          user.AccessToken,
		Role:                                 user.Role,
		FavoriteLandProperties:               user.FavoriteLandProperties,
		FavoriteStands:                       user.FavoriteStands,
		FavoriteCommercialForSaleProperties:  user.FavoriteCommercialForSaleProperties,
		FavoriteCommercialRentalProperties:   user.FavoriteCommercialRentalProperties,
		FavoriteResidentialForSaleProperties: user.FavoriteResidentialForSaleProperties,
		FavoriteResidentialRentalProperties:  user.FavoriteCommercialRentalProperties,
		ProfilePicture: userDtos.UserProfilePictureResponseDTO{
			Id:          user.ProfilePicture.Id,
			UserId:      user.ProfilePicture.UserId,
			Uri:         user.ProfilePicture.Uri,
			Name:        user.ProfilePicture.Name,
			FileType:    user.ProfilePicture.FileType,
			Size:        user.ProfilePicture.Size,
			ContentType: user.ProfilePicture.ContentType,
		},
		Location: userDtos.LocationDTO{
			UserId:      user.Location.UserId,
			Lat:         user.Location.Lat,
			Lon:         user.Location.Lon,
			City:        user.Location.City,
			County:      user.Location.County,
			Country:     user.Location.Country,
			CountryCode: user.Location.CountryCode,
			Surburb:     user.Location.Surburb,
			Id:          user.Location.Id,
			Boundingbox: user.Location.Boundingbox,
			DisplayName: user.Location.DisplayName,
			Province:    user.Location.Province,
		},
	}
}
