// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package server

import (
	"context"

	"api/app/core/cmd/rpc/internal/logic/userservice"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	core.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) SignUp(ctx context.Context, in *core.SignUpReq) (*core.SignUpResp, error) {
	l := userservicelogic.NewSignUpLogic(ctx, s.svcCtx)
	return l.SignUp(in)
}

func (s *UserServiceServer) SignIn(ctx context.Context, in *core.SignInReq) (*core.SignInResp, error) {
	l := userservicelogic.NewSignInLogic(ctx, s.svcCtx)
	return l.SignIn(in)
}

func (s *UserServiceServer) GetUserInfo(ctx context.Context, in *core.GetUserInfoReq) (*core.GetUserInfoResp, error) {
	l := userservicelogic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserServiceServer) GetUserFriendProfile(ctx context.Context, in *core.GetUserFriendProfileReq) (*core.GetUserFriendProfileResp, error) {
	l := userservicelogic.NewGetUserFriendProfileLogic(ctx, s.svcCtx)
	return l.GetUserFriendProfile(in)
}

func (s *UserServiceServer) UpdateUserInfo(ctx context.Context, in *core.UpdateUserInfoReq) (*core.UpdateUserInfoResp, error) {
	l := userservicelogic.NewUpdateUserInfoLogic(ctx, s.svcCtx)
	return l.UpdateUserInfo(in)
}

func (s *UserServiceServer) UpdateUserStatus(ctx context.Context, in *core.UpdateUserStatusReq) (*core.UpdateUserStatusResp, error) {
	l := userservicelogic.NewUpdateUserStatusLogic(ctx, s.svcCtx)
	return l.UpdateUserStatus(in)
}

func (s *UserServiceServer) UploadUserAvatar(ctx context.Context, in *core.UploadUserAvatarReq) (*core.UploadUserAvatarResp, error) {
	l := userservicelogic.NewUploadUserAvatarLogic(ctx, s.svcCtx)
	return l.UploadUserAvatar(in)
}

func (s *UserServiceServer) UploadUserCover(ctx context.Context, in *core.UploadUserCoverReq) (*core.UploadUserCoverResp, error) {
	l := userservicelogic.NewUploadUserCoverLogic(ctx, s.svcCtx)
	return l.UploadUserCover(in)
}

func (s *UserServiceServer) SearchUser(ctx context.Context, in *core.SearchUserReq) (*core.SearchUserResp, error) {
	l := userservicelogic.NewSearchUserLogic(ctx, s.svcCtx)
	return l.SearchUser(in)
}

func (s *UserServiceServer) AddUserSticker(ctx context.Context, in *core.AddStickerReq) (*core.AddStickerResp, error) {
	l := userservicelogic.NewAddUserStickerLogic(ctx, s.svcCtx)
	return l.AddUserSticker(in)
}

func (s *UserServiceServer) DeleteUserSticker(ctx context.Context, in *core.DeleteStickerReq) (*core.DeleteStickerResp, error) {
	l := userservicelogic.NewDeleteUserStickerLogic(ctx, s.svcCtx)
	return l.DeleteUserSticker(in)
}

func (s *UserServiceServer) IsStickerExist(ctx context.Context, in *core.IsStickerExistReq) (*core.IsStickerExistResp, error) {
	l := userservicelogic.NewIsStickerExistLogic(ctx, s.svcCtx)
	return l.IsStickerExist(in)
}

func (s *UserServiceServer) GetUserStickers(ctx context.Context, in *core.GetUserStickerReq) (*core.GetUserStickerResp, error) {
	l := userservicelogic.NewGetUserStickersLogic(ctx, s.svcCtx)
	return l.GetUserStickers(in)
}
