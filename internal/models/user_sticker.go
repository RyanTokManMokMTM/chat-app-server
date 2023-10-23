package models

type UserSticker struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	UserId    uint   `gorm:"not null"`
	StickerId string `gorm:"type:varchar(64);not null"`

	User    User    `gorm:"foreignKey:UserId"`
	Sticker Sticker `gorm:"foreignKey:StickerId;references:Uuid"`
	CommonField
}

func (us *UserSticker) TableName() string {
	return "users_stickers"
}
