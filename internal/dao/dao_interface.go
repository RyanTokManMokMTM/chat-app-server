package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
	socket_message "github.com/ryantokmanmok/chat-app-server/socket-proto"
)

type DAOInterface interface {
	InsertOneUser(ctx context.Context, name, email, password string) (*models.UserModel, error)
	FindOneUser(ctx context.Context, id uint) (*models.UserModel, error)
	FindOneUserByEmail(ctx context.Context, email string) (*models.UserModel, error)
	UpdateUserProfile(ctx context.Context, id uint, name string) error
	UpdateUserAvatar(ctx context.Context, id uint, avatarName string) error
	FindUsers(ctx context.Context, query string) ([]*models.UserModel, error)
	CountUserAvailableStory(ctx context.Context, userID uint) (int64, error)

	InsertOneFriend(ctx context.Context, userID, friendID uint) error
	FindOneFriend(ctx context.Context, userID, friendID uint) error
	DeleteOneFriend(ctx context.Context, userID, friendID uint) error
	GetUserFriendList(ctx context.Context, userID uint) ([]*models.UserFriend, error)

	InsertOneGroup(ctx context.Context, groupName string, userID uint) (*models.Group, error)
	FindOneGroup(ctx context.Context, groupID uint) (*models.Group, error)
	FindOneGroupByUUID(ctx context.Context, groupUUID string) (*models.Group, error)
	DeleteOneGroup(ctx context.Context, groupID uint) error
	UpdateOneGroup(ctx context.Context, groupID uint, groupName string) error
	UpdateOneGroupAvatar(ctx context.Context, groupID uint, avatarName string) error
	GetUserGroups(ctx context.Context, userID uint) ([]*models.GroupMember, error)

	InsertOneGroupMember(ctx context.Context, groupID, userID uint) error
	FindOneGroupMember(ctx context.Context, groupID, userID uint) (*models.GroupMember, error)
	FindOneGroupMembers(ctx context.Context, groupID uint) ([]*models.GroupMember, error)
	DeleteGroupMember(ctx context.Context, groupID, userID uint) error
	DeleteAllGroupMembers(ctx context.Context, groupID uint) error
	GetGroupMembers(ctx context.Context, groupID uint) ([]*models.GroupMember, error)

	InsertOneMessage(ctx context.Context, message *socket_message.Message)
	FindOneMessage(ctx context.Context, messageID uint) (*models.Message, error)
	DeleteOneMessage(ctx context.Context, messageID uint) error
	GetMessage(ctx context.Context, from, to, messageType uint) ([]*models.Message, error)

	InsertOneStory(ctx context.Context, userID uint, mediaPath string) (uint, error)
	FindOneStory(ctx context.Context, storyID uint) (*models.StoryModel, error)
	FindOneUserStory(ctx context.Context, storyID, userID uint) (*models.StoryModel, error)
	GetUserStories(ctx context.Context, userID uint) ([]uint, error)
	GetActiveUsers(ctx context.Context, userID uint) ([]*models.UserModel, error)
	DeleteStories(ctx context.Context, storyID uint) error
}
