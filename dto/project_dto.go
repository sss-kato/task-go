package dto

import "time"

type ProjectDto struct {
	ID        uint
	Name      string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ud ProjectDto) TableName() string {
	return "projects"
}
