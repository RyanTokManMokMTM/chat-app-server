package models

type CreateUserGroupDTO struct {
	UserId  uint
	GroupId uint
}

type UpdateUserGroupDTO struct {
	UserId  uint
	GroupId uint
}

const userGroupTableName = "users_groups"

type UserGroup struct {
	Base
	GroupID uint `gorm:"not uniqueIndex;index;comment:'belong to which group ID'"`
	UserID  uint `gorm:"not uniqueIndex;index;comment:'who belong to this group'"`

	MemberInfo User  `gorm:"foreignKey:UserID"`
	GroupInfo  Group `gorm:"foreignKey:GroupID"`
}

func (ug *UserGroup) TableName() string {
	return userGroupTableName
}
