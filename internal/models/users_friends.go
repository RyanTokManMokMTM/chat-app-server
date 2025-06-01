package models

type CreateUserFriendDTO struct {
	UserId   uint
	FriendId uint
}

type UpdateUserFriendDTO struct {
	UserId   uint
	FriendId uint
}

type UserFriend struct {
	//user A added user B ,but it doesn't mean user B has added userA ?
	Base
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserId   uint `gorm:"not null;index"`
	FriendID uint `gorm:"not null;index"`

	//UserInfo   User `gorm:"foreignKey:UserId"`
	FriendInfo User `gorm:"foreignKey:FriendID"`
}

func (uf *UserFriend) TableName() string {
	return "users_friends"
}
