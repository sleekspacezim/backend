package repositories

import (
	constants "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	"math"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := c.DefaultQuery("page", "1")
		pageLimitStr := c.DefaultQuery("pageLimit", "10")
		page := generalUtilities.ConvertStringToInt(pageStr)
		pageLimit := generalUtilities.ConvertStringToInt(pageLimitStr)
		if page <= 0 {
			page = 1
		}
		if pageLimit <= 0 {
			pageLimit = constants.PageLimit
		}

		dbClone := db.Session(&gorm.Session{})
		var total int64
		dbClone.Count(&total)
		totalPages := int(math.Ceil(float64(total) / float64(pageLimit)))

		c.Set("totalPages", totalPages)
		c.Set("count", total)

		offSet := (page - 1) * pageLimit
		return db.Offset(offSet).Limit(pageLimit)
	}
}
