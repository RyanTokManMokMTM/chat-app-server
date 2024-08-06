// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package stickerservice

import (
	"context"

	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddFriendReq             = core.AddFriendReq
	AddFriendResp            = core.AddFriendResp
	AddStickerReq            = core.AddStickerReq
	AddStickerResp           = core.AddStickerResp
	AddStoryReq              = core.AddStoryReq
	AddStoryResp             = core.AddStoryResp
	CountUserGroupReq        = core.CountUserGroupReq
	CountUserGroupResp       = core.CountUserGroupResp
	CreateGroupReq           = core.CreateGroupReq
	CreateGroupResp          = core.CreateGroupResp
	CreateStickerGroupReq    = core.CreateStickerGroupReq
	CreateStickerGroupResp   = core.CreateStickerGroupResp
	CreateStoryLikeReq       = core.CreateStoryLikeReq
	CreateStoryLikeResp      = core.CreateStoryLikeResp
	DeleteFriendReq          = core.DeleteFriendReq
	DeleteFriendResp         = core.DeleteFriendResp
	DeleteGroupReq           = core.DeleteGroupReq
	DeleteGroupResp          = core.DeleteGroupResp
	DeleteStickerReq         = core.DeleteStickerReq
	DeleteStickerResp        = core.DeleteStickerResp
	DeleteStoryLikeReq       = core.DeleteStoryLikeReq
	DeleteStoryLikeResp      = core.DeleteStoryLikeResp
	DeleteStoryReq           = core.DeleteStoryReq
	DeleteStoryResp          = core.DeleteStoryResp
	FriendInfo               = core.FriendInfo
	FriendStory              = core.FriendStory
	FullGroupInfo            = core.FullGroupInfo
	GetActiveStoryReq        = core.GetActiveStoryReq
	GetActiveStoryResp       = core.GetActiveStoryResp
	GetFriendInfoReq         = core.GetFriendInfoReq
	GetFriendInfoResp        = core.GetFriendInfoResp
	GetFriendListReq         = core.GetFriendListReq
	GetFriendListResp        = core.GetFriendListResp
	GetGroupInfoByUUIDReq    = core.GetGroupInfoByUUIDReq
	GetGroupInfoByUUIDResp   = core.GetGroupInfoByUUIDResp
	GetGroupMembersReq       = core.GetGroupMembersReq
	GetGroupMembersResp      = core.GetGroupMembersResp
	GetStickerInfoReq        = core.GetStickerInfoReq
	GetStickerInfoResp       = core.GetStickerInfoResp
	GetStickerListReq        = core.GetStickerListReq
	GetStickerListResp       = core.GetStickerListResp
	GetStickerResourcesReq   = core.GetStickerResourcesReq
	GetStickerResourcesResp  = core.GetStickerResourcesResp
	GetStoryInfoByIdRep      = core.GetStoryInfoByIdRep
	GetStoryInfoByIdResp     = core.GetStoryInfoByIdResp
	GetStorySeenListReq      = core.GetStorySeenListReq
	GetStorySeenListResp     = core.GetStorySeenListResp
	GetUserFriendProfileReq  = core.GetUserFriendProfileReq
	GetUserFriendProfileResp = core.GetUserFriendProfileResp
	GetUserGroupReq          = core.GetUserGroupReq
	GetUserGroupResp         = core.GetUserGroupResp
	GetUserInfoReq           = core.GetUserInfoReq
	GetUserInfoResp          = core.GetUserInfoResp
	GetUserStickerReq        = core.GetUserStickerReq
	GetUserStickerResp       = core.GetUserStickerResp
	GetUserStoryReq          = core.GetUserStoryReq
	GetUserStoryResp         = core.GetUserStoryResp
	GroupInfo                = core.GroupInfo
	GroupMemberInfo          = core.GroupMemberInfo
	IsStickerExistReq        = core.IsStickerExistReq
	IsStickerExistResp       = core.IsStickerExistResp
	JoinGroupReq             = core.JoinGroupReq
	JoinGroupResp            = core.JoinGroupResp
	LeaveGroupReq            = core.LeaveGroupReq
	LeaveGroupResp           = core.LeaveGroupResp
	PageableInfo             = core.PageableInfo
	SearchGroupReq           = core.SearchGroupReq
	SearchGroupResp          = core.SearchGroupResp
	SearchUserReq            = core.SearchUserReq
	SearchUserResp           = core.SearchUserResp
	SearchUserRespResult     = core.SearchUserRespResult
	SignInReq                = core.SignInReq
	SignInResp               = core.SignInResp
	SignUpReq                = core.SignUpReq
	SignUpResp               = core.SignUpResp
	StickerData              = core.StickerData
	StickerFileMap           = core.StickerFileMap
	StickerInfo              = core.StickerInfo
	StoryInfo                = core.StoryInfo
	StorySeenInfo            = core.StorySeenInfo
	StorySeenUserBasicInfo   = core.StorySeenUserBasicInfo
	UpdateGroupInfoReq       = core.UpdateGroupInfoReq
	UpdateGroupInfoResp      = core.UpdateGroupInfoResp
	UpdateStorySeenReq       = core.UpdateStorySeenReq
	UpdateStorySeenResp      = core.UpdateStorySeenResp
	UpdateUserInfoReq        = core.UpdateUserInfoReq
	UpdateUserInfoResp       = core.UpdateUserInfoResp
	UpdateUserStatusReq      = core.UpdateUserStatusReq
	UpdateUserStatusResp     = core.UpdateUserStatusResp
	UploadGroupAvatarReq     = core.UploadGroupAvatarReq
	UploadGroupAvatarResp    = core.UploadGroupAvatarResp
	UploadUserAvatarReq      = core.UploadUserAvatarReq
	UploadUserAvatarResp     = core.UploadUserAvatarResp
	UploadUserCoverReq       = core.UploadUserCoverReq
	UploadUserCoverResp      = core.UploadUserCoverResp
	UserInfo                 = core.UserInfo

	StickerService interface {
		CreateStickerGroup(ctx context.Context, in *CreateStickerGroupReq, opts ...grpc.CallOption) (*CreateStickerGroupResp, error)
		GetStickerGroupResources(ctx context.Context, in *GetStickerResourcesReq, opts ...grpc.CallOption) (*GetStickerResourcesResp, error)
		GetStickerGroupInfo(ctx context.Context, in *GetStickerInfoReq, opts ...grpc.CallOption) (*GetStickerInfoResp, error)
		GetStickerGroupList(ctx context.Context, in *GetStickerListReq, opts ...grpc.CallOption) (*GetStickerListResp, error)
	}

	defaultStickerService struct {
		cli zrpc.Client
	}
)

func NewStickerService(cli zrpc.Client) StickerService {
	return &defaultStickerService{
		cli: cli,
	}
}

func (m *defaultStickerService) CreateStickerGroup(ctx context.Context, in *CreateStickerGroupReq, opts ...grpc.CallOption) (*CreateStickerGroupResp, error) {
	client := core.NewStickerServiceClient(m.cli.Conn())
	return client.CreateStickerGroup(ctx, in, opts...)
}

func (m *defaultStickerService) GetStickerGroupResources(ctx context.Context, in *GetStickerResourcesReq, opts ...grpc.CallOption) (*GetStickerResourcesResp, error) {
	client := core.NewStickerServiceClient(m.cli.Conn())
	return client.GetStickerGroupResources(ctx, in, opts...)
}

func (m *defaultStickerService) GetStickerGroupInfo(ctx context.Context, in *GetStickerInfoReq, opts ...grpc.CallOption) (*GetStickerInfoResp, error) {
	client := core.NewStickerServiceClient(m.cli.Conn())
	return client.GetStickerGroupInfo(ctx, in, opts...)
}

func (m *defaultStickerService) GetStickerGroupList(ctx context.Context, in *GetStickerListReq, opts ...grpc.CallOption) (*GetStickerListResp, error) {
	client := core.NewStickerServiceClient(m.cli.Conn())
	return client.GetStickerGroupList(ctx, in, opts...)
}
