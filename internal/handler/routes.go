// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "github.com/ryantokmanmokmtm/chat-app-server/internal/handler/file"
	friend "github.com/ryantokmanmokmtm/chat-app-server/internal/handler/friend"
	group "github.com/ryantokmanmokmtm/chat-app-server/internal/handler/group"
	health "github.com/ryantokmanmokmtm/chat-app-server/internal/handler/health"
	message "github.com/ryantokmanmokmtm/chat-app-server/internal/handler/message"
	sticker "github.com/ryantokmanmokmtm/chat-app-server/internal/handler/sticker"
	story "github.com/ryantokmanmokmtm/chat-app-server/internal/handler/story"
	user "github.com/ryantokmanmokmtm/chat-app-server/internal/handler/user"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: health.HealthCheckHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/signup",
				Handler: user.UserSignUpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/signin",
				Handler: user.UserSignInHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/profile",
				Handler: user.GetUserFriendProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/user/info",
				Handler: user.UpdateUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/user/status",
				Handler: user.UpdateUserStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/avatar",
				Handler: user.UploadUserAvatarHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/cover",
				Handler: user.UploadUserCoverHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/search",
				Handler: user.SearchUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/sticker",
				Handler: user.AddUserStickerHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/user/sticker",
				Handler: user.DeleteUserStickerHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/sticker/:sticker_id",
				Handler: user.IsStickerExistHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/sticker/list",
				Handler: user.GetUserStickersHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/friend",
				Handler: friend.AddFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/friend",
				Handler: friend.DeleteFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/friends",
				Handler: friend.GetFriendListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/friend/:uuid",
				Handler: friend.GetFriendInformationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/group",
				Handler: group.CreateGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/group/join/:group_id",
				Handler: group.JoinGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/group/leave/:group_id",
				Handler: group.LeaveGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/group",
				Handler: group.DeleteGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/members/:group_id",
				Handler: group.GetGroupMembersHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/group/avatar/:group_id",
				Handler: group.UploadGroupAvatarHandler(serverCtx),
			},
			{
				Method:  http.MethodPatch,
				Path:    "/group/update",
				Handler: group.UpdateGroupInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group",
				Handler: group.GetUserGroupsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/search",
				Handler: group.SearchGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/info/uuid/:uuid",
				Handler: group.GetGroupInfoByUUIDHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/group/count",
				Handler: group.CountUserGroupHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/message",
				Handler: message.GetMessagesHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/message",
				Handler: message.DeleteMessageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/file/image/upload",
				Handler: file.UploadImageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload",
				Handler: file.UploadFileHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/story",
				Handler: story.AddStoryHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/story",
				Handler: story.DeleteStoryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/stories/:user_id",
				Handler: story.GetUserStoriesByUserIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/stories/active",
				Handler: story.GetActiveStoriesHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/story/seen",
				Handler: story.UpdateStorySeenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/story/like",
				Handler: story.CreateStoryLikeHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/story/like",
				Handler: story.DeleteStoryLikeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/story/:story_id",
				Handler: story.GetStoryInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/story/seen/:story_id",
				Handler: story.GetStorySeenListInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/sticker",
				Handler: sticker.CreateStickerGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/sticker/resources/:sticker_group_uuid",
				Handler: sticker.GetStickerGroupResourcesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/sticker/:sticker_uuid",
				Handler: sticker.GetStickerGroupInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/sticker/list",
				Handler: sticker.GetStickerGroupListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
