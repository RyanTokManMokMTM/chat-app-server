package logic

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/config"
	mock_dao "github.com/ryantokmanmokmtm/chat-app-server/internal/dao/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserSignUp(t *testing.T) {
	mock := gomock.NewController(t)
	_ = mock_dao.NewMockStore(mock)
	defer mock.Finish()

	var c config.Config
	configBytes := []byte(`{
		"Mode" : "test", 
		"Path" : "./resources",
		"Auth":{
				"AccessSecret": "2BNVfmf0WtyX1HQmzYG5rOKLzlHBEPRX729pZ0gpxujnaikoRRCF78T8fKDNTLWy",
				"AccessExpire": 86400
		},
		"Salt":"W4tiDEeWlwxlRPYYRRMhJ63piS1ochvMymwfVdumittPoSxhkHNnZLe6m12C4v15",
		"MaxBytes": 524288000
	

	}`)

	err := json.Unmarshal(configBytes, &c)
	assert.Nil(t, err)
	//
	//api := "/api/v1/user/signup"
	//user := models.UserModel{
	//	Email:    "",
	//	Password: "",
	//}
	//type testCasesStruct struct {
	//	TestName string
	//	mockTest func(store *mock_dao.MockStore)
	//	response func(t *testing.T, recorder *httptest.ResponseRecorder)
	//}
	//
	////MARK: SignUp Request
	//reqEmail := "admin@admin.com"
	//reqPassword := "admin12345"
	//mockdb.EXPECT().FindOneUserByEmail(gomock.Any(), reqEmail).Times(1).Return(nil, gorm.ErrRecordNotFound)
	//mockdb.EXPECT().InsertOneUser(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(&models.UserModel{
	//	Email: reqEmail,
	//}, nil)

}

func TestUserSignIn(t *testing.T) {

}

func TestUploadUserAvatar(t *testing.T) {

}

func TestUploadUserCover(t *testing.T) {

}

func TestUpdateUserStatusMessage(t *testing.T) {

}

func TestUpdateUserInfo(t *testing.T) {

}

func TestSearchUser(t *testing.T) {

}

func TestGetUserInfo(t *testing.T) {

}

func TestGetUserFriendProfile(t *testing.T) {

}
