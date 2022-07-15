package handlers

import (
	. "dimacros/best-price-seeker/app/desired-product"

	"gorm.io/gorm"
)

func CreateDesiredProduct(command *CreateDesiredProductCommand, db *gorm.DB) DesiredProduct {
	desiredProduct := CreateDesiredProductFromCommand(command)

	db.Create(&desiredProduct)

	return desiredProduct
}
