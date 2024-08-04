// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package groupservice

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

	GroupService interface {
		CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error)
		JoinGroup(ctx context.Context, in *JoinGroupReq, opts ...grpc.CallOption) (*JoinGroupResp, error)
		LeaveGroup(ctx context.Context, in *LeaveGroupReq, opts ...grpc.CallOption) (*LeaveGroupResp, error)
		DeleteGroup(ctx context.Context, in *DeleteGroupReq, opts ...grpc.CallOption) (*DeleteGroupResp, error)
		GetGroupMembers(ctx context.Context, in *GetGroupMembersReq, opts ...grpc.CallOption) (*GetGroupMembersResp, error)
		UploadGroupAvatar(ctx context.Context, opts ...grpc.CallOption) (core.GroupService_UploadGroupAvatarClient, error)
		UpdateGroupInfo(ctx context.Context, in *UpdateGroupInfoReq, opts ...grpc.CallOption) (*UpdateGroupInfoResp, error)
		GetUserGroups(ctx context.Context, in *GetUserGroupReq, opts ...grpc.CallOption) (*GetUserGroupResp, error)
		SearchGroup(ctx context.Context, in *SearchGroupReq, opts ...grpc.CallOption) (*SearchGroupResp, error)
		GetGroupInfoByUUID(ctx context.Context, in *GetGroupInfoByUUIDReq, opts ...grpc.CallOption) (*GetGroupInfoByUUIDResp, error)
		CountUserGroup(ctx context.Context, in *CountUserGroupReq, opts ...grpc.CallOption) (*CountUserGroupResp, error)
	}

	defaultGroupService struct {
		cli zrpc.Client
	}
)

func NewGroupService(cli zrpc.Client) GroupService {
	return &defaultGroupService{
		cli: cli,
	}
}

func (m *defaultGroupService) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.CreateGroup(ctx, in, opts...)
}

func (m *defaultGroupService) JoinGroup(ctx context.Context, in *JoinGroupReq, opts ...grpc.CallOption) (*JoinGroupResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.JoinGroup(ctx, in, opts...)
}

func (m *defaultGroupService) LeaveGroup(ctx context.Context, in *LeaveGroupReq, opts ...grpc.CallOption) (*LeaveGroupResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.LeaveGroup(ctx, in, opts...)
}

func (m *defaultGroupService) DeleteGroup(ctx context.Context, in *DeleteGroupReq, opts ...grpc.CallOption) (*DeleteGroupResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.DeleteGroup(ctx, in, opts...)
}

func (m *defaultGroupService) GetGroupMembers(ctx context.Context, in *GetGroupMembersReq, opts ...grpc.CallOption) (*GetGroupMembersResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.GetGroupMembers(ctx, in, opts...)
}

func (m *defaultGroupService) UploadGroupAvatar(ctx context.Context, opts ...grpc.CallOption) (core.GroupService_UploadGroupAvatarClient, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.UploadGroupAvatar(ctx, opts...)
}

func (m *defaultGroupService) UpdateGroupInfo(ctx context.Context, in *UpdateGroupInfoReq, opts ...grpc.CallOption) (*UpdateGroupInfoResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.UpdateGroupInfo(ctx, in, opts...)
}

func (m *defaultGroupService) GetUserGroups(ctx context.Context, in *GetUserGroupReq, opts ...grpc.CallOption) (*GetUserGroupResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.GetUserGroups(ctx, in, opts...)
}

func (m *defaultGroupService) SearchGroup(ctx context.Context, in *SearchGroupReq, opts ...grpc.CallOption) (*SearchGroupResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.SearchGroup(ctx, in, opts...)
}

func (m *defaultGroupService) GetGroupInfoByUUID(ctx context.Context, in *GetGroupInfoByUUIDReq, opts ...grpc.CallOption) (*GetGroupInfoByUUIDResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.GetGroupInfoByUUID(ctx, in, opts...)
}

func (m *defaultGroupService) CountUserGroup(ctx context.Context, in *CountUserGroupReq, opts ...grpc.CallOption) (*CountUserGroupResp, error) {
	client := core.NewGroupServiceClient(m.cli.Conn())
	return client.CountUserGroup(ctx, in, opts...)
}
