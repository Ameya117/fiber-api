package models

import (
	"time"
)

type User struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	FirstName  string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
