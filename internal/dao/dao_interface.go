package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
)

type Store interface {
	InsertOneUser(ctx context.Context, name, email, password string) (*models.User, error)
	FindOneUser(ctx context.Context, id uint) (*models.User, error)
	FindOneUserByEmail(ctx context.Context, email string) (*models.User, error)
	FindOneUserByUUID(ctx context.Context, uuid string) (*models.User, error)
	UpdateUserProfile(ctx context.Context, id uint, name string) error
	UpdateUserStatusMessage(ctx context.Context, id uint, message string) error
	UpdateUserAvatar(ctx context.Context, id uint, avatarPath string) error
	UpdateUserCover(ctx context.Context, id uint, coverPath string) error
	FindUsers(ctx context.Context, query string) ([]*models.User, error)
	CountUserAvailableStory(ctx context.Context, userID uint) (int64, error)

	InsertOneFriend(ctx context.Context, userID, friendID uint) error
	FindOneFriend(ctx context.Context, userID, friendID uint) (*models.User, error)
	DeleteOneFriend(ctx context.Context, userID, friendID uint) error
	GetUserFriendListByPageSize(ctx context.Context, userID uint, pageOffset, PageLimit int) ([]*models.UserFriend, error)
	CountUserFriend(ctx context.Context, userID uint) (int64, error)

	InsertOneGroup(ctx context.Context, groupName, avatar string, userID uint) (*models.Group, error)
	FindOneGroup(ctx context.Context, groupID uint) (*models.Group, error)
	FindOneGroupByUUID(ctx context.Context, groupUUID string) (*models.Group, error)
	DeleteOneGroup(ctx context.Context, groupID uint) error
	UpdateOneGroup(ctx context.Context, groupID uint, groupName string, groupDesc string) error
	UpdateOneGroupAvatar(ctx context.Context, groupID uint, avatarName string) error
	GetUserGroups(ctx context.Context, userID uint, pageOffset, pageLimit int) ([]*models.UserGroup, error)
	CountUserGroups(ctx context.Context, userID uint) int64
	SearchGroup(ctx context.Context, query string) ([]*models.Group, error)

	InsertOneGroupMember(ctx context.Context, groupID, userID uint) error
	FindOneGroupMember(ctx context.Context, groupID, userID uint) (*models.UserGroup, error)
	FindOneGroupMembers(ctx context.Context, groupID uint) ([]*models.UserGroup, error)
	FindOneGroupMembersByPage(ctx context.Context, groupID uint, pageOffset, pageLimit int) ([]*models.UserGroup, error)
	DeleteGroupMember(ctx context.Context, groupID, userID uint) error
	DeleteAllGroupMembers(ctx context.Context, groupID uint) error
	GetGroupMembers(ctx context.Context, groupID uint, pageOffset, pageLimit int) ([]*models.UserGroup, error)
	CountGroupMembers(ctx context.Context, groupID uint) (int64, error)

	InsertOneMessage(ctx context.Context, message *socket_message.Message)
	FindOneMessage(ctx context.Context, messageID uint) (*models.Message, error)
	DeleteOneMessage(ctx context.Context, messageID uint) error
	CountMessage(ctx context.Context, messageType, id uint) (int64, error)
	GetMessage(ctx context.Context, from, to, messageType uint, pageLimit int, latestId uint) ([]*models.Message, error)

	InsertOneStory(ctx context.Context, userID uint, mediaPath string) (uint, error)
	FindOneStory(ctx context.Context, storyID uint) (*models.StoryModel, error)
	FindOneUserStory(ctx context.Context, storyID, userID uint) (*models.StoryModel, error)
	GetUserStories(ctx context.Context, userID uint) ([]uint, error)
	GetUserStoriesByTimeStamp(ctx context.Context, userID uint, timeStamp int64) ([]uint, error)
	GetFriendActiveStories(ctx context.Context, userID uint, pageOffset, pageLimit int) ([]*models.StoriesWithLatestStoryTime, error)
	GetFriendActiveStoriesByTimeStamp(ctx context.Context, userID uint, pageOffset, pageLimit int, timeStamp int64) ([]*models.StoriesWithLatestStoryTime, error)
	DeleteStories(ctx context.Context, storyID uint) error
	CountActiveStoryByTimeStamp(ctx context.Context, userId uint, timeStamp int64) (int64, error)

	InsertOneUserStorySeen(ctx context.Context, userID, friendId, storyId uint) error
	FindOneUserStorySeen(ctx context.Context, userID, friendId uint) (*models.UserStorySeen, error)
	UpdateOneUserStorySeen(ctx context.Context, Id, storyId uint) error
	DeleteOneUserStorySeen(ctx context.Context, ID uint) error

	InsertOneUserStoryLike(ctx context.Context, userID, storyId uint) error
	FindOneUserStoryLike(ctx context.Context, userID, storyId uint) (*models.UserStoryLikes, error)
	CountStoryLikes(ctx context.Context, storyId uint) (int64, error)
	DeleteOneUserStoryLike(ctx context.Context, ID uint) error
}
