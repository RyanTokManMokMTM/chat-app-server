package models

type CreateUserGroupDTO struct {
	UserId  uint
	GroupId uint
}

type UpdateUserGroupDTO struct {
	UserId  uint
	GroupId uint
}

type UserGroup struct {
	Base
	ID      uint `gorm:"primaryKey;autoIncrement"`
	GroupId uint `gorm:"not uniqueIndex;index;comment:'belong to which group Id'"`
	UserId  uint `gorm:"not uniqueIndex;index;comment:'who belong to this group'"`

	MemberInfo User  `gorm:"foreignKey:UserId"`
	GroupInfo  Group `gorm:"foreignKey:GroupId"`
}

func (ug *UserGroup) TableName() string {
	return "users_groups"
}
