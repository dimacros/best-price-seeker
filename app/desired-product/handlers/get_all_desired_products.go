package handlers

import (
	. "dimacros/best-price-seeker/app/desired-product"

	"gorm.io/gorm"
)

func GetAllDesiredProducts(db *gorm.DB) []DesiredProduct {
	var desiredProducts []DesiredProduct

	db.Find(&desiredProducts)

	return desiredProducts
}
