package admin

import (
	"net/http"

	managerRepo "SleekSpace/repositories/manager"

	"github.com/gin-gonic/gin"
)

func GetAllManagersContacts(c *gin.Context) {
	numbers := managerRepo.GetAllManagersContacts()
	c.JSON(http.StatusOK, gin.H{
		"response": numbers,
	})
}

func GetAllManagers(c *gin.Context) {
	managers := managerRepo.GetAllManagers()
	c.JSON(http.StatusOK, gin.H{
		"response": managers,
	})
}

func GetAllManagersProfilePictures(c *gin.Context) {
	pictures := managerRepo.GetAllManagersProfilePictures()
	c.JSON(http.StatusOK, gin.H{
		"response": pictures,
	})
}
