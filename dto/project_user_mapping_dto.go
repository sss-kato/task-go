package dto

type Project_User_MappingDto struct {
	ID        uint
	ProjectID uint
	UserID    uint
}

func (pumd Project_User_MappingDto) TableName() string {
	return "project_user_mappings"
}
