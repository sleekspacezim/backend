package scopes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PropertyStructureTypeFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	structureType := c.DefaultQuery("type", "")
	if structureType != "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("type= ?", structureType)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func CurrencyFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	currency := c.DefaultQuery("currency", "")
	if currency != "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("currency= ?", currency)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func RentFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	rentMin := c.DefaultQuery("rent_min", "0")
	rentMax := c.DefaultQuery("rent_max", "0")
	if rentMax != "0" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(
				"rent_amount BETWEEN ? AND ?",
				rentMin,
				rentMax,
			)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func PriceFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	priceMin := c.DefaultQuery("price_min", "0")
	priceMax := c.DefaultQuery("price_max", "0")
	if priceMax != "0" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("price BETWEEN ? AND ?", priceMin, priceMax)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func NumberOfRoomsFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	totalRooms := c.DefaultQuery("number_of_rooms", "")
	if totalRooms == "8+" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("number_of_rooms > ?", 7)
		}
	} else if totalRooms != "" && totalRooms != "8+" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("number_of_rooms = ?", totalRooms)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func NumberOfRoomsToLetFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	numberOfRoomsToLet := c.DefaultQuery("number_of_rooms_to_rent", "")
	if numberOfRoomsToLet == "8+" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("number_of_rooms_to_let > ?", 7)
		}
	} else if numberOfRoomsToLet != "" && numberOfRoomsToLet != "8+" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("number_of_rooms_to_let = ?", numberOfRoomsToLet)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func BedroomsFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	numberOfRoomsBedrooms := c.DefaultQuery("bedrooms", "")
	if numberOfRoomsBedrooms == "8+" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("bedrooms > ?", 7)
		}
	} else if numberOfRoomsBedrooms != "" && numberOfRoomsBedrooms != "8+" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("bedrooms = ?", numberOfRoomsBedrooms)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func BathroomsFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	numberOfRoomsBathrooms := c.DefaultQuery("bathrooms", "")
	if numberOfRoomsBathrooms == "8+" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("bathrooms > ?", 7)
		}
	} else if numberOfRoomsBathrooms != "" && numberOfRoomsBathrooms != "8+" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("bathrooms = ?", numberOfRoomsBathrooms)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}

func PropertySizeFilter(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	dimensions := c.DefaultQuery("size_dimensions", "")
	sizeNumber := c.DefaultQuery("size_number", "")
	if dimensions != "" && sizeNumber != "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("size_number = ? AND size_dimensions = ?", sizeNumber, dimensions)
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db
		}
	}
}
