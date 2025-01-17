package models

import "time"

type SourceProduct struct {
	ID           int64     `gorm:"primaryKey" json:"id"`
	ProductName  string    `json:"product_name"`
	Qty          int64     `json:"qty"`
	SellingPrice float64   `json:"selling_price"`
	PromoPrice   float64   `json:"promo_price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
