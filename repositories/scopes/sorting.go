package scopes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SortProperties(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	sortingField := c.DefaultQuery("sortBy", "newest")
	if sortingField == "newest" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC, id DESC")
		}
	}
	if sortingField == "oldest" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at ASC, id ASC")
		}
	}
	if sortingField == "rent_low_to_high" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("rent_amount ASC, id DESC")
		}
	}
	if sortingField == "rent_high_to_low" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("rent_amount DESC, id DESC")
		}
	}
	if sortingField == "price_low_to_high" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("price ASC, id DESC")
		}
	}
	if sortingField == "price_high_to_low" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("price DESC, id DESC")
		}
	}
	if sortingField == "rooms_to_rent_low_to_high" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("is_full_house ASC, number_of_rooms_to_let ASC, id DESC")
		}
	}
	if sortingField == "rooms_to_rent_high_to_low" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("is_full_house DESC, number_of_rooms_to_let DESC, id DESC")
		}
	}
	if sortingField == "total_rooms_low_to_high" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("total_number_of_rooms ASC, id DESC")
		}
	}
	if sortingField == "total_rooms_high_to_low" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("total_number_of_rooms DESC, id DESC")
		}
	}
	if sortingField == "size_low_to_high" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("size_dimensions DESC, size_number ASC, id DESC")
		}
	}
	if sortingField == "size_high_to_low" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("size_dimensions ASC, size_number DESC, id DESC")
		}
	}
	if sortingField == "year_built_oldest_to_new" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("year_built ASC, id DESC")
		}
	} else {
		return func(db *gorm.DB) *gorm.DB {
			return db.Order("year_built DESC, id DESC")
		}
	}
}
