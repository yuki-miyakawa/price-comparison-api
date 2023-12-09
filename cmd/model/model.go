package model

import "time"

type Product struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	UnitPrice int       `json:"unit_price"`
	Unit      string    `json:"unit"`
	ShopName  string    `json:"shop_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
