package admin

import (
	"net/http"

	reportRepo "SleekSpace/repositories/property/report"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
)

func GetAllPropertiesReports(c *gin.Context) {
	reports := reportRepo.GetAllPropertiesReports()
	if reports == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "reports does not exist."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ProcessedPropertyReportsListToResponse(reports)})
}
