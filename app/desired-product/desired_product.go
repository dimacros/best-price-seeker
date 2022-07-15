package desiredproduct

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ProductLink struct {
	gorm.Model
	DesiredProductID uint
	Cost             float32
	Link             string
}

type DesiredProduct struct {
	gorm.Model
	Name          string
	ProductLinks  []ProductLink
	BestPrice     float32
	StartingPrice float32
}

type CreateProductLinkCommand struct {
	Cost float32 `json:"cost"`
	Link string  `json:"link"`
}

type CreateDesiredProductCommand struct {
	Name         string                     `json:"description"`
	ProductLinks []CreateProductLinkCommand `json:"product_links"`
}

func CreateDesiredProductFromCommand(command *CreateDesiredProductCommand) DesiredProduct {
	var ProductLinks []ProductLink

	for _, value := range command.ProductLinks {
		ProductLinks = append(ProductLinks, ProductLink{Cost: value.Cost, Link: value.Link})
	}

	return DesiredProduct{
		Name:          command.Name,
		BestPrice:     100,
		StartingPrice: 100,
		ProductLinks:  ProductLinks,
	}
}
