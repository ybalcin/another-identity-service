package account

import (
	"errors"
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/location"
)

func TestAddNewUser(t *testing.T) {
	users := []*user{}

	repo := new(mockRepository)
	repo.InsertNewUserFn = func(user *user) *errRepository {
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
	repo.InsertNewUserFn = func(user *user) *errRepository {
		return &errRepository{
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
