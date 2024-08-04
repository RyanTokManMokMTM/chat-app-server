// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package server

import (
	"context"

	"api/app/core/cmd/rpc/internal/logic/groupservice"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
)

type GroupServiceServer struct {
	svcCtx *svc.ServiceContext
	core.UnimplementedGroupServiceServer
}

func NewGroupServiceServer(svcCtx *svc.ServiceContext) *GroupServiceServer {
	return &GroupServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *GroupServiceServer) CreateGroup(ctx context.Context, in *core.CreateGroupReq) (*core.CreateGroupResp, error) {
	l := groupservicelogic.NewCreateGroupLogic(ctx, s.svcCtx)
	return l.CreateGroup(in)
}

func (s *GroupServiceServer) JoinGroup(ctx context.Context, in *core.JoinGroupReq) (*core.JoinGroupResp, error) {
	l := groupservicelogic.NewJoinGroupLogic(ctx, s.svcCtx)
	return l.JoinGroup(in)
}

func (s *GroupServiceServer) LeaveGroup(ctx context.Context, in *core.LeaveGroupReq) (*core.LeaveGroupResp, error) {
	l := groupservicelogic.NewLeaveGroupLogic(ctx, s.svcCtx)
	return l.LeaveGroup(in)
}

func (s *GroupServiceServer) DeleteGroup(ctx context.Context, in *core.DeleteGroupReq) (*core.DeleteGroupResp, error) {
	l := groupservicelogic.NewDeleteGroupLogic(ctx, s.svcCtx)
	return l.DeleteGroup(in)
}

func (s *GroupServiceServer) GetGroupMembers(ctx context.Context, in *core.GetGroupMembersReq) (*core.GetGroupMembersResp, error) {
	l := groupservicelogic.NewGetGroupMembersLogic(ctx, s.svcCtx)
	return l.GetGroupMembers(in)
}

func (s *GroupServiceServer) UploadGroupAvatar(stream core.GroupService_UploadGroupAvatarServer) error {
	l := groupservicelogic.NewUploadGroupAvatarLogic(stream.Context(), s.svcCtx)
	return l.UploadGroupAvatar(stream)
}

func (s *GroupServiceServer) UpdateGroupInfo(ctx context.Context, in *core.UpdateGroupInfoReq) (*core.UpdateGroupInfoResp, error) {
	l := groupservicelogic.NewUpdateGroupInfoLogic(ctx, s.svcCtx)
	return l.UpdateGroupInfo(in)
}

func (s *GroupServiceServer) GetUserGroups(ctx context.Context, in *core.GetUserGroupReq) (*core.GetUserGroupResp, error) {
	l := groupservicelogic.NewGetUserGroupsLogic(ctx, s.svcCtx)
	return l.GetUserGroups(in)
}

func (s *GroupServiceServer) SearchGroup(ctx context.Context, in *core.SearchGroupReq) (*core.SearchGroupResp, error) {
	l := groupservicelogic.NewSearchGroupLogic(ctx, s.svcCtx)
	return l.SearchGroup(in)
}

func (s *GroupServiceServer) GetGroupInfoByUUID(ctx context.Context, in *core.GetGroupInfoByUUIDReq) (*core.GetGroupInfoByUUIDResp, error) {
	l := groupservicelogic.NewGetGroupInfoByUUIDLogic(ctx, s.svcCtx)
	return l.GetGroupInfoByUUID(in)
}

func (s *GroupServiceServer) CountUserGroup(ctx context.Context, in *core.CountUserGroupReq) (*core.CountUserGroupResp, error) {
	l := groupservicelogic.NewCountUserGroupLogic(ctx, s.svcCtx)
	return l.CountUserGroup(in)
}