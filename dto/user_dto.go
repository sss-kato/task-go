package dto

import "time"

type UserDto struct {
	ID         uint
	Name       string
	Password   string
	Mailadress string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (ud UserDto) TableName() string {
	return "users"
}
