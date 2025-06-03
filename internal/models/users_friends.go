package models

type CreateUserFriendDTO struct {
	UserId   uint
	FriendId uint
}

type UpdateUserFriendDTO struct {
	UserId   uint
	FriendId uint
}

const userFriendTableName = "users_friends"

type UserFriend struct {
	//user A added user B ,but it doesn't mean user B has added userA ?
	Base
	UserID   uint `gorm:"not null;index"`
	FriendID uint `gorm:"not null;index"`

	FriendInfo User `gorm:"foreignKey:FriendID"`
}

func (uf *UserFriend) TableName() string {
	return userFriendTableName
}
