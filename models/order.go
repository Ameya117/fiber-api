package models

import "time"

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	ProductRefer int       `json:"product_id"`
	Product      Product   `gorm:"foreignKey:ProductRefer"`
	UserRefer    int       `json:"user_id"`
	User         User      `gorm:"foreignKey:UserRefer"`
	CreatedAt    time.Time `json:"created_at"`
}
