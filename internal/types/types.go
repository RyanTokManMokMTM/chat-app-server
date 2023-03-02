// Code generated by goctl. DO NOT EDIT.
package types

type HealthCheckReq struct {
}

type HealthCheckResp struct {
	Resp string `json:"resp"`
}

type CommonUserInfo struct {
	ID       uint   `json:"user_id"`
	Uuid     string `json:"uuid"`
	NickName string `json:"name"`
	Avatar   string `json:"avatar"`
}

type SignUpReq struct {
	Email    string `json:"email" validate:"email,min=8,max=32"`
	Name     string `json:"name" validate:"min=8,max=16"`
	Password string `json:"password" validate:"min=8,max=32"`
}

type SignUpResp struct {
	Code        uint   `json:"code"`
	Token       string `json:"token"`
	ExpiredTime uint   `json:"expired_time"`
}

type SignInReq struct {
	Email    string `json:"email" validate:"email,min=8,max=32"`
	Password string `json:"password" validate:"min=8,max=32"`
}

type SignInResp struct {
	Code        uint   `json:"code"`
	Token       string `json:"token"`
	ExpiredTime uint   `json:"expired_time"`
}

type GetUserInfoReq struct {
	UserID uint `path:"user_id"`
}

type GetUserInfoResp struct {
	Code   uint   `json:"code"`
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type UpdateUserInfoReq struct {
	Name string `json:"name" validate:"min=8,max=32"`
}

type UpdateUserInfoResp struct {
	Code uint `json:"code"`
}

type UploadUserAvatarReq struct {
}

type UploadUserAvatarResp struct {
	Code uint   `json:"code"`
	Path string `json:"path"`
}

type AddFriendReq struct {
	UserID uint `json:"user_id"`
}

type AddFriendResp struct {
	Code uint `json:"code"`
}

type DeleteFriendReq struct {
	UserID uint `json:"user_id"`
}

type DeleteFriendResp struct {
	Code uint `json:"code"`
}

type GetFriendListReq struct {
}

type GetFriendListResp struct {
	FriendList []FriendInfo `json:"friend_list"`
}

type FriendInfo struct {
	ID       uint   `json:"user_id"`
	UUID     string `json:"uuid"`
	NickName string `json:"name"`
	Avatar   string `json:"avatar"`
}

type CreateGroupReq struct {
	GroupName string `json:"group_name"`
}

type CreateGroupResp struct {
	Code    uint `json:"code"`
	GroupID uint `json:"group_id"`
}

type JoinGroupReq struct {
	GroupID uint `path:"group_id"`
}

type JoinGroupResp struct {
	Code uint `json:"code"`
}

type LeaveGroupReq struct {
	GroupID uint `path:"group_id"`
}

type LeaveGroupResp struct {
	Code uint `json:"code"`
}

type DeleteGroupReq struct {
	GroupID uint `json:"group_id"`
}

type DeleteGroupResp struct {
	Code uint `json:"code"`
}

type GetGroupMembersReq struct {
	GroupID uint `path:"group_id"`
}

type GetGroupMembersResp struct {
	Code       uint              `json:"code"`
	MemberList []GroupMemberInfo `json:"member_list"`
}

type UpdateGroupInfoReq struct {
	GroupName string `json:"group_name"`
}

type UpdateGroupInfoResp struct {
	Code uint `json:"code"`
}

type UploadGroupAvatarReq struct {
}

type UploadGroupAvatarResp struct {
	Code uint `json:"code"`
}

type GroupMemberInfo struct {
	CommonUserInfo
	IsGroupLead bool `json:"is_group_lead"`
}

type GetMessagesReq struct {
	ID          uint `json:"id"` //can be a user id or a group id
	MessageType uint `json:"message_type"`
	FriendID    uint `json:"friend_id"` //only for message type = 1
}

type GetMessagesResp struct {
	Code     uint          `json:"code"`
	Messages []MessageUser `json:"message"`
}

type DeleteMessageReq struct {
	MesssageID uint `json:"msg_id"`
}

type DeleteMessageResp struct {
	Code uint `json:"code"`
}

type MessageUser struct {
	MessageID   uint   `json:"message_id"`
	FromID      uint   `json:"from_id"`
	ToID        uint   `json:"to_id"`
	Content     string `json:"cotent"`
	ContentType uint   `json:"content_type"`
	MessageType uint   `json:"message_type"`
	CreatedAt   uint   `json:"create_at"`
}
