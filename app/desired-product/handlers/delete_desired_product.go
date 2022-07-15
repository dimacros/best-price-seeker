package handlers

import (
	. "dimacros/best-price-seeker/app/desired-product"

	"gorm.io/gorm"
)

func DeleteDesiredProduct(ID int, db *gorm.DB) {
	db.Delete(&DesiredProduct{}, ID)
}
