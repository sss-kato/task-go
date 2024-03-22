package model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null"`
}
