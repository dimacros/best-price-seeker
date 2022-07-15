package main

import (
	. "dimacros/best-price-seeker/app/desired-product"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&ProductLink{}, &DesiredProduct{})
}
