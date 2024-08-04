// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ryantokmanmokmtm/chat-app-router/internal/dao (interfaces: Store)

// Package mock_dao is a generated GoMock package.
package mock_dao

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
)

// MockStore is a mock of Store interfaces.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CountGroupMembers mocks base method.
func (m *MockStore) CountGroupMembers(arg0 context.Context, arg1 uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountGroupMembers", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountGroupMembers indicates an expected call of CountGroupMembers.
func (mr *MockStoreMockRecorder) CountGroupMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountGroupMembers", reflect.TypeOf((*MockStore)(nil).CountGroupMembers), arg0, arg1)
}

// CountUserAvailableStory mocks base method.
func (m *MockStore) CountUserAvailableStory(arg0 context.Context, arg1 uint) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUserAvailableStory", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUserAvailableStory indicates an expected call of CountUserAvailableStory.
func (mr *MockStoreMockRecorder) CountUserAvailableStory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUserAvailableStory", reflect.TypeOf((*MockStore)(nil).CountUserAvailableStory), arg0, arg1)
}

// DeleteAllGroupMembers mocks base method.
func (m *MockStore) DeleteAllGroupMembers(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllGroupMembers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllGroupMembers indicates an expected call of DeleteAllGroupMembers.
func (mr *MockStoreMockRecorder) DeleteAllGroupMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllGroupMembers", reflect.TypeOf((*MockStore)(nil).DeleteAllGroupMembers), arg0, arg1)
}

// DeleteGroupMember mocks base method.
func (m *MockStore) DeleteGroupMember(arg0 context.Context, arg1, arg2 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroupMember indicates an expected call of DeleteGroupMember.
func (mr *MockStoreMockRecorder) DeleteGroupMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupMember", reflect.TypeOf((*MockStore)(nil).DeleteGroupMember), arg0, arg1, arg2)
}

// DeleteOneFriend mocks base method.
func (m *MockStore) DeleteOneFriend(arg0 context.Context, arg1, arg2 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOneFriend", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOneFriend indicates an expected call of DeleteOneFriend.
func (mr *MockStoreMockRecorder) DeleteOneFriend(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOneFriend", reflect.TypeOf((*MockStore)(nil).DeleteOneFriend), arg0, arg1, arg2)
}

// DeleteOneGroup mocks base method.
func (m *MockStore) DeleteOneGroup(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOneGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOneGroup indicates an expected call of DeleteOneGroup.
func (mr *MockStoreMockRecorder) DeleteOneGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOneGroup", reflect.TypeOf((*MockStore)(nil).DeleteOneGroup), arg0, arg1)
}

// DeleteOneMessage mocks base method.
func (m *MockStore) DeleteOneMessage(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOneMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOneMessage indicates an expected call of DeleteOneMessage.
func (mr *MockStoreMockRecorder) DeleteOneMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOneMessage", reflect.TypeOf((*MockStore)(nil).DeleteOneMessage), arg0, arg1)
}

// DeleteStories mocks base method.
func (m *MockStore) DeleteStories(arg0 context.Context, arg1 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStories", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStories indicates an expected call of DeleteStories.
func (mr *MockStoreMockRecorder) DeleteStories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStories", reflect.TypeOf((*MockStore)(nil).DeleteStories), arg0, arg1)
}

// FindOneFriend mocks base method.
func (m *MockStore) FindOneFriend(arg0 context.Context, arg1, arg2 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneFriend", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindOneFriend indicates an expected call of FindOneFriend.
func (mr *MockStoreMockRecorder) FindOneFriend(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneFriend", reflect.TypeOf((*MockStore)(nil).FindOneFriend), arg0, arg1, arg2)
}

// FindOneGroup mocks base method.
func (m *MockStore) FindOneGroup(arg0 context.Context, arg1 uint) (*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneGroup", arg0, arg1)
	ret0, _ := ret[0].(*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneGroup indicates an expected call of FindOneGroup.
func (mr *MockStoreMockRecorder) FindOneGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneGroup", reflect.TypeOf((*MockStore)(nil).FindOneGroup), arg0, arg1)
}

// FindOneGroupByUUID mocks base method.
func (m *MockStore) FindOneGroupByUUID(arg0 context.Context, arg1 string) (*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneGroupByUUID", arg0, arg1)
	ret0, _ := ret[0].(*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneGroupByUUID indicates an expected call of FindOneGroupByUUID.
func (mr *MockStoreMockRecorder) FindOneGroupByUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneGroupByUUID", reflect.TypeOf((*MockStore)(nil).FindOneGroupByUUID), arg0, arg1)
}

// FindOneGroupMember mocks base method.
func (m *MockStore) FindOneGroupMember(arg0 context.Context, arg1, arg2 uint) (*models.UserGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneGroupMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.UserGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneGroupMember indicates an expected call of FindOneGroupMember.
func (mr *MockStoreMockRecorder) FindOneGroupMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneGroupMember", reflect.TypeOf((*MockStore)(nil).FindOneGroupMember), arg0, arg1, arg2)
}

// FindOneGroupMembers mocks base method.
func (m *MockStore) FindOneGroupMembers(arg0 context.Context, arg1 uint) ([]*models.UserGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneGroupMembers", arg0, arg1)
	ret0, _ := ret[0].([]*models.UserGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneGroupMembers indicates an expected call of FindOneGroupMembers.
func (mr *MockStoreMockRecorder) FindOneGroupMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneGroupMembers", reflect.TypeOf((*MockStore)(nil).FindOneGroupMembers), arg0, arg1)
}

// FindOneMessage mocks base method.
func (m *MockStore) FindOneMessage(arg0 context.Context, arg1 uint) (*models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneMessage", arg0, arg1)
	ret0, _ := ret[0].(*models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneMessage indicates an expected call of FindOneMessage.
func (mr *MockStoreMockRecorder) FindOneMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneMessage", reflect.TypeOf((*MockStore)(nil).FindOneMessage), arg0, arg1)
}

// FindOneStory mocks base method.
func (m *MockStore) FindOneStory(arg0 context.Context, arg1 uint) (*models.StoryModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneStory", arg0, arg1)
	ret0, _ := ret[0].(*models.StoryModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneStory indicates an expected call of FindOneStory.
func (mr *MockStoreMockRecorder) FindOneStory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneStory", reflect.TypeOf((*MockStore)(nil).FindOneStory), arg0, arg1)
}

// FindOneUser mocks base method.
func (m *MockStore) FindOneUser(arg0 context.Context, arg1 uint) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneUser indicates an expected call of FindOneUser.
func (mr *MockStoreMockRecorder) FindOneUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneUser", reflect.TypeOf((*MockStore)(nil).FindOneUser), arg0, arg1)
}

// FindOneUserByEmail mocks base method.
func (m *MockStore) FindOneUserByEmail(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneUserByEmail", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneUserByEmail indicates an expected call of FindOneUserByEmail.
func (mr *MockStoreMockRecorder) FindOneUserByEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneUserByEmail", reflect.TypeOf((*MockStore)(nil).FindOneUserByEmail), arg0, arg1)
}

// FindOneUserByUUID mocks base method.
func (m *MockStore) FindOneUserByUUID(arg0 context.Context, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneUserByUUID", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneUserByUUID indicates an expected call of FindOneUserByUUID.
func (mr *MockStoreMockRecorder) FindOneUserByUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneUserByUUID", reflect.TypeOf((*MockStore)(nil).FindOneUserByUUID), arg0, arg1)
}

// FindOneUserStory mocks base method.
func (m *MockStore) FindOneUserStory(arg0 context.Context, arg1, arg2 uint) (*models.StoryModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneUserStory", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.StoryModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneUserStory indicates an expected call of FindOneUserStory.
func (mr *MockStoreMockRecorder) FindOneUserStory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneUserStory", reflect.TypeOf((*MockStore)(nil).FindOneUserStory), arg0, arg1, arg2)
}

// FindUsers mocks base method.
func (m *MockStore) FindUsers(arg0 context.Context, arg1 string) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUsers", arg0, arg1)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUsers indicates an expected call of FindUsers.
func (mr *MockStoreMockRecorder) FindUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUsers", reflect.TypeOf((*MockStore)(nil).FindUsers), arg0, arg1)
}

// GetActiveUsers mocks base method.
func (m *MockStore) GetActiveUsers(arg0 context.Context, arg1 uint) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveUsers", arg0, arg1)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveUsers indicates an expected call of GetActiveUsers.
func (mr *MockStoreMockRecorder) GetActiveUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveUsers", reflect.TypeOf((*MockStore)(nil).GetActiveUsers), arg0, arg1)
}

// GetGroupMembers mocks base method.
func (m *MockStore) GetGroupMembers(arg0 context.Context, arg1 uint) ([]*models.UserGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMembers", arg0, arg1)
	ret0, _ := ret[0].([]*models.UserGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMembers indicates an expected call of GetGroupMembers.
func (mr *MockStoreMockRecorder) GetGroupMembers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMembers", reflect.TypeOf((*MockStore)(nil).GetGroupMembers), arg0, arg1)
}

// GetMessage mocks base method.
func (m *MockStore) GetMessage(arg0 context.Context, arg1, arg2, arg3 uint) ([]*models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockStoreMockRecorder) GetMessage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockStore)(nil).GetMessage), arg0, arg1, arg2, arg3)
}

// GetUserFriendList mocks base method.
func (m *MockStore) GetUserFriendList(arg0 context.Context, arg1 uint) ([]*models.UserFriend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFriendList", arg0, arg1)
	ret0, _ := ret[0].([]*models.UserFriend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFriendList indicates an expected call of GetUserFriendList.
func (mr *MockStoreMockRecorder) GetUserFriendList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFriendList", reflect.TypeOf((*MockStore)(nil).GetUserFriendList), arg0, arg1)
}

// GetUserGroups mocks base method.
func (m *MockStore) GetUserGroups(arg0 context.Context, arg1 uint) ([]*models.UserGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserGroups", arg0, arg1)
	ret0, _ := ret[0].([]*models.UserGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserGroups indicates an expected call of GetUserGroups.
func (mr *MockStoreMockRecorder) GetUserGroups(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserGroups", reflect.TypeOf((*MockStore)(nil).GetUserGroups), arg0, arg1)
}

// GetUserStories mocks base method.
func (m *MockStore) GetUserStories(arg0 context.Context, arg1 uint) ([]uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserStories", arg0, arg1)
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserStories indicates an expected call of GetUserStories.
func (mr *MockStoreMockRecorder) GetUserStories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStories", reflect.TypeOf((*MockStore)(nil).GetUserStories), arg0, arg1)
}

// InsertOneFriend mocks base method.
func (m *MockStore) InsertOneFriend(arg0 context.Context, arg1, arg2 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOneFriend", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOneFriend indicates an expected call of InsertOneFriend.
func (mr *MockStoreMockRecorder) InsertOneFriend(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOneFriend", reflect.TypeOf((*MockStore)(nil).InsertOneFriend), arg0, arg1, arg2)
}

// InsertOneGroup mocks base method.
func (m *MockStore) InsertOneGroup(arg0 context.Context, arg1, arg2 string, arg3 uint) (*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOneGroup", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOneGroup indicates an expected call of InsertOneGroup.
func (mr *MockStoreMockRecorder) InsertOneGroup(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOneGroup", reflect.TypeOf((*MockStore)(nil).InsertOneGroup), arg0, arg1, arg2, arg3)
}

// InsertOneGroupMember mocks base method.
func (m *MockStore) InsertOneGroupMember(arg0 context.Context, arg1, arg2 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOneGroupMember", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOneGroupMember indicates an expected call of InsertOneGroupMember.
func (mr *MockStoreMockRecorder) InsertOneGroupMember(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOneGroupMember", reflect.TypeOf((*MockStore)(nil).InsertOneGroupMember), arg0, arg1, arg2)
}

// InsertOneMessage mocks base method.
func (m *MockStore) InsertOneMessage(arg0 context.Context, arg1 *socket_message.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InsertOneMessage", arg0, arg1)
}

// InsertOneMessage indicates an expected call of InsertOneMessage.
func (mr *MockStoreMockRecorder) InsertOneMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOneMessage", reflect.TypeOf((*MockStore)(nil).InsertOneMessage), arg0, arg1)
}

// InsertOneStory mocks base method.
func (m *MockStore) InsertOneStory(arg0 context.Context, arg1 uint, arg2 string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOneStory", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOneStory indicates an expected call of InsertOneStory.
func (mr *MockStoreMockRecorder) InsertOneStory(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOneStory", reflect.TypeOf((*MockStore)(nil).InsertOneStory), arg0, arg1, arg2)
}

// InsertOneUser mocks base method.
func (m *MockStore) InsertOneUser(arg0 context.Context, arg1, arg2, arg3 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOneUser", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOneUser indicates an expected call of InsertOneUser.
func (mr *MockStoreMockRecorder) InsertOneUser(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOneUser", reflect.TypeOf((*MockStore)(nil).InsertOneUser), arg0, arg1, arg2, arg3)
}

// SearchGroup mocks base method.
func (m *MockStore) SearchGroup(arg0 context.Context, arg1 string) ([]*models.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchGroup", arg0, arg1)
	ret0, _ := ret[0].([]*models.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchGroup indicates an expected call of SearchGroup.
func (mr *MockStoreMockRecorder) SearchGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchGroup", reflect.TypeOf((*MockStore)(nil).SearchGroup), arg0, arg1)
}

// UpdateOneGroup mocks base method.
func (m *MockStore) UpdateOneGroup(arg0 context.Context, arg1 uint, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOneGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOneGroup indicates an expected call of UpdateOneGroup.
func (mr *MockStoreMockRecorder) UpdateOneGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOneGroup", reflect.TypeOf((*MockStore)(nil).UpdateOneGroup), arg0, arg1, arg2)
}

// UpdateOneGroupAvatar mocks base method.
func (m *MockStore) UpdateOneGroupAvatar(arg0 context.Context, arg1 uint, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOneGroupAvatar", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOneGroupAvatar indicates an expected call of UpdateOneGroupAvatar.
func (mr *MockStoreMockRecorder) UpdateOneGroupAvatar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOneGroupAvatar", reflect.TypeOf((*MockStore)(nil).UpdateOneGroupAvatar), arg0, arg1, arg2)
}

// UpdateUserAvatar mocks base method.
func (m *MockStore) UpdateUserAvatar(arg0 context.Context, arg1 uint, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserAvatar", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserAvatar indicates an expected call of UpdateUserAvatar.
func (mr *MockStoreMockRecorder) UpdateUserAvatar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAvatar", reflect.TypeOf((*MockStore)(nil).UpdateUserAvatar), arg0, arg1, arg2)
}

// UpdateUserCover mocks base method.
func (m *MockStore) UpdateUserCover(arg0 context.Context, arg1 uint, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserCover", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserCover indicates an expected call of UpdateUserCover.
func (mr *MockStoreMockRecorder) UpdateUserCover(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserCover", reflect.TypeOf((*MockStore)(nil).UpdateUserCover), arg0, arg1, arg2)
}

// UpdateUserProfile mocks base method.
func (m *MockStore) UpdateUserProfile(arg0 context.Context, arg1 uint, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserProfile indicates an expected call of UpdateUserProfile.
func (mr *MockStoreMockRecorder) UpdateUserProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserProfile", reflect.TypeOf((*MockStore)(nil).UpdateUserProfile), arg0, arg1, arg2)
}

// UpdateUserStatusMessage mocks base method.
func (m *MockStore) UpdateUserStatusMessage(arg0 context.Context, arg1 uint, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserStatusMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserStatusMessage indicates an expected call of UpdateUserStatusMessage.
func (mr *MockStoreMockRecorder) UpdateUserStatusMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserStatusMessage", reflect.TypeOf((*MockStore)(nil).UpdateUserStatusMessage), arg0, arg1, arg2)
}