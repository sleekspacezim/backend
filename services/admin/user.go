package admin

import (
	"net/http"

	userRepo "SleekSpace/repositories/user"
	generalUtilities "SleekSpace/utilities/funcs/general"

	"github.com/gin-gonic/gin"
)

func GetAllUsersLocations(c *gin.Context) {
	codes := userRepo.GetAllUsersLocations()
	c.JSON(http.StatusOK, gin.H{
		"response": codes,
	})
}

func GetAllUsers(c *gin.Context) {
	users := userRepo.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"response": users,
	})
}

func GetAllUsersProfilePictures(c *gin.Context) {
	pictures := userRepo.GetAllUsersProfilePictures()
	c.JSON(http.StatusOK, gin.H{
		"response": pictures,
	})
}

func GetVerificationCodeById(c *gin.Context) {
	id := c.Param("id")
	code := userRepo.GetVerificationCodeById(id)
	c.JSON(http.StatusOK, gin.H{
		"response": code,
	})
}

func GetAllVerificationCodes(c *gin.Context) {
	codes := userRepo.AllVerificationCodes()
	c.JSON(http.StatusOK, gin.H{
		"response": codes,
	})
}

func DeleteVerificationCode(c *gin.Context) {
	id := c.Param("id")
	isDeleted := userRepo.DeleteVerficationCode(generalUtilities.ConvertStringToInt(id))
	if isDeleted {
		c.JSON(http.StatusOK, gin.H{
			"response": "code deleted",
		})
	}
}

func GetUserContacts(c *gin.Context) {
	numbers := userRepo.GetAllUsersContactNumbers()
	c.JSON(http.StatusOK, gin.H{
		"response": numbers,
	})
}
