package manager

import (
	managerDtos "SleekSpace/dtos/manager"
	managerModels "SleekSpace/models/manager"
)

func AddManagerIdToContacts(contacts []managerModels.ManagerContactNumber, managerId int) []managerModels.ManagerContactNumber {
	newContacts := []managerModels.ManagerContactNumber{}
	for i := 0; i < len(contacts); i++ {
		contact := managerModels.ManagerContactNumber{
			Id:           contacts[i].Id,
			CountryAbbrv: contacts[i].CountryAbbrv,
			CountryCode:  contacts[i].CountryCode,
			Number:       contacts[i].Number,
			Type:         contacts[i].Type,
			ManagerId:    managerId,
		}
		newContacts = append(newContacts, contact)
	}
	return newContacts
}

func processedManagerContactNumbers(contactNumbers []managerModels.ManagerContactNumber) []managerDtos.ManagerContactNumberDTO {
	contacts := []managerDtos.ManagerContactNumberDTO{}
	for i := 0; i < len(contactNumbers); i++ {
		contact := managerDtos.ManagerContactNumberDTO{
			Id:           contactNumbers[i].Id,
			CountryAbbrv: contactNumbers[i].CountryAbbrv,
			CountryCode:  contactNumbers[i].CountryCode,
			Number:       contactNumbers[i].Number,
			Type:         contactNumbers[i].Type,
			ManagerId:    contactNumbers[i].ManagerId,
		}
		contacts = append(contacts, contact)
	}
	return contacts
}

func ManagerResponse(manager *managerModels.Manager) managerDtos.ManagerResponseDTO {
	return managerDtos.ManagerResponseDTO{
		Id:     manager.Id,
		UserId: manager.UserId,
		Email:  manager.Email,
		Name:   manager.Name,
		ProfilePicture: managerDtos.ManagerProfilePictureResponseDTO{
			Id:          manager.ProfilePicture.Id,
			ManagerId:   manager.ProfilePicture.ManagerId,
			Uri:         manager.ProfilePicture.Uri,
			Name:        manager.ProfilePicture.Name,
			FileType:    manager.ProfilePicture.FileType,
			Size:        manager.ProfilePicture.Size,
			ContentType: manager.ProfilePicture.ContentType,
		},
		Contacts: processedManagerContactNumbers(manager.ManagerContactNumbers),
	}
}
