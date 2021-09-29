package account

import (
	"errors"
	"github.com/ybalcin/another-identity-service/store"
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/location"
)

func TestAccountService_AddNewUser(t *testing.T) {
	store.MgoConfig = store.MongoConfig{
		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
		Database: "another-identity-accountService-store",
	}
	store.InitMongo()

	var users []*user

	repo := new(mockRepository)
	repo.InsertNewUserFn = func(user *user) *common.FriendlyError {
		users = append(users, user)
		return nil
	}
	service := NewAccountService(repo)
	service.AddNewUser("jon", "doe", "johndoe", "johndoe@gmail.com", "123456", time.Now(), "5344444", true, &location.Address{
		Country: "tr",
		City:    "ist",
		County:  "kçekmece",
	})

	for _, u := range users {
		if u == nil || u.Firstname != "jon" {
			t.FailNow()
		}
	}

}

func TestAccountService_AddNewUserErrRepository(t *testing.T) {
	repo := new(mockRepository)
	repo.InsertNewUserFn = func(user *user) *common.FriendlyError {
		return &common.FriendlyError{
			Message:        "test message",
			InnerException: errors.New("test error"),
		}
	}

	service := NewAccountService(repo)
	errService := service.AddNewUser("jon", "doe", "johndoe", "johndoe@gmail.com", "123456", time.Now(), "5344444", true, &location.Address{
		Country: "tr",
		City:    "ist",
		County:  "kçekmece",
	})

	if errService == nil || errService.FriendlyMessage != err_insert_new_user {
		t.Errorf("TestAddNewUserErrRepository fail!")
	}
}

func TestAccountService_AddRoleToUser(t *testing.T) {
	userRepository := NewUserRepository()
	service := NewAccountService(userRepository)

	if err := service.AddRoleToUser("guest", "60f70c937c6ebf50d5464366"); err != nil {
		t.Errorf(err.Message)
	}
}

func TestAccountService_GetUserList_ShouldReturn_Error_When_Fail(t *testing.T) {
	repo := new(mockRepository)
	repo.GetUserListFn = func() ([]*user, *common.FriendlyError) {
		return nil, &common.FriendlyError{}
	}

	service := NewAccountService(repo)
	_, err := service.GetUserList()
	if err == nil {
		t.Errorf("Should be error!")
	}
}

func TestAccountService_GetUserList_ShouldReturn_NilError_When_Success(t *testing.T) {
	repo := new(mockRepository)
	repo.GetUserListFn = func() ([]*user, *common.FriendlyError) {
		return []*user{}, nil
	}

	service := NewAccountService(repo)
	_, err := service.GetUserList()
	if err != nil {
		t.Errorf("Error should be nil!")
	}
}
