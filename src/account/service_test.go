package account

import (
	"errors"
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/location"
	"github.com/ybalcin/another-identity-service/store"
)

func TestAddNewUser(t *testing.T) {
	users := []*user{}

	repo := new(mockRepository)
	repo.InsertNewUserFn = func(user *user) *common.FriendlyError {
		users = append(users, user)
		return nil
	}
	service := NewService(repo)
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

func TestAddNewUserErrRepository(t *testing.T) {
	repo := new(mockRepository)
	repo.InsertNewUserFn = func(user *user) *common.FriendlyError {
		return &common.FriendlyError{
			Message:        "test message",
			InnerException: errors.New("test error"),
		}
	}

	service := NewService(repo)
	errService := service.AddNewUser("jon", "doe", "johndoe", "johndoe@gmail.com", "123456", time.Now(), "5344444", true, &location.Address{
		Country: "tr",
		City:    "ist",
		County:  "kçekmece",
	})

	if errService == nil || errService.FriendlyMessage != insertNewUserError {
		t.Errorf("TestAddNewUserErrRepository fail!")
	}
}

func TestAddRoleToUser(t *testing.T) {
	store.MgoConfig = store.MongoConfig{
		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
		Database: "another-identity-service-store",
	}
	store.InitMongo()

	user_repository := NewUserRepository()
	service := NewService(user_repository)

	if err := service.AddRoleToUser("guest", "60f70c937c6ebf50d5464366"); err != nil {
		t.Errorf(err.Message)
	}
}
