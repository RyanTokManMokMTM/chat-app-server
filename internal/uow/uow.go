package unitofwork

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/repository"
	"gorm.io/gorm"
)

type UOW struct {
	engine               *gorm.DB
	userRepo             repository.IUserRepo[models.User]
	groupRepo            repository.IGroupRepo[models.Group]
	messageRepo          repository.IMessageRepo[models.Message]
	stickerRepo          repository.IStickerRepo[models.Sticker]
	stickerResourcesRepo repository.IStickerResourcesRepo[models.StickerResource]
	storyRepo            repository.IStoryRepo[models.StoryModel]
	userFriendsRepo      repository.IUserFriendsRepo[models.UserFriend]
	userGroupRepo        repository.IUserGroupRepo[models.UserGroup]
	userStoryLikesRepo   repository.IUserStoryLikesRepo[models.UserStoryLikes]
	userStorySeenRepo    repository.IUserStorySeenRepo[models.UserStorySeen]
}

func NewUOW(engine *gorm.DB) *UOW {
	return &UOW{
		engine:               engine,
		userRepo:             repository.NewUserRepo(engine),
		groupRepo:            repository.NewGroupRepo(engine),
		messageRepo:          repository.NewMessageRepo(engine),
		stickerRepo:          repository.NewStickerRepo(engine),
		stickerResourcesRepo: repository.NewStickerResourcesRepo(engine),
		storyRepo:            repository.NewStoryRepo(engine),
		userFriendsRepo:      repository.NewUserFriendsRepo(engine),
		userGroupRepo:        repository.NewUserGroupRepo(engine),
		userStoryLikesRepo:   repository.NewUserStoryLikesRepo(engine),
		userStorySeenRepo:    repository.NewUserStorySeenRepo(engine),
	}
}

func (u *UOW) UserRepo() repository.IUserRepo[models.User] {
	return u.userRepo
}

func (u *UOW) GroupRepo() repository.IGroupRepo[models.Group] {
	return u.groupRepo
}

func (u *UOW) MessageRepo() repository.IMessageRepo[models.Message] {
	return u.messageRepo
}

func (u *UOW) StickerRepo() repository.IStickerRepo[models.Sticker] {
	return u.stickerRepo
}

func (u *UOW) StickerResourcesRepo() repository.IStickerResourcesRepo[models.StickerResource] {
	return u.stickerResourcesRepo
}

func (u *UOW) StoryRepo() repository.IStoryRepo[models.StoryModel] {
	return u.storyRepo
}

func (u *UOW) UserFriendsRepo() repository.IUserFriendsRepo[models.UserFriend] {
	return u.userFriendsRepo
}

func (u *UOW) UserGroupRepo() repository.IUserGroupRepo[models.UserGroup] {
	return u.userGroupRepo
}

func (u *UOW) UserStoryLikesRepo() repository.IUserStoryLikesRepo[models.UserStoryLikes] {
	return u.userStoryLikesRepo
}

func (u *UOW) UserStorySeenRepo() repository.IUserStorySeenRepo[models.UserStorySeen] {
	return u.userStorySeenRepo
}

func (u *UOW) WithTransaction(ctx context.Context, fn func(context.Context, *UOW) error) error {
	return u.engine.Transaction(func(tx *gorm.DB) error {
		return fn(ctx, NewUOW(tx))
	})
}
