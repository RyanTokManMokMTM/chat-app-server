package models

type CreateGroupDTO struct {
	GroupName   string
	GroupAvatar string
	GroupDesc   string
	GroupLead   uint
}

type UpdateGroupDTO struct {
	GroupName   string
	GroupAvatar string
	GroupDesc   string
}

const groupTableColName = "groups_info"

type Group struct {
	Base
	GroupName   string `gorm:"types:varchar(64);not null"`
	GroupAvatar string
	GroupDesc   string
	GroupLead   uint `gorm:"not null;index"`

	LeadInfo  User   `gorm:"foreignKey:GroupLead;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UsersInfo []User `gorm:"many2many:users_groups;foreignKey:ID;joinForeignKey:GroupID"`
}

func (g *Group) TableName() string {
	return groupTableColName
}
