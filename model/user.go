package model

import "time"

type User struct {
	ID         uint
	Name       string
	Password   string
	Mailadress string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
