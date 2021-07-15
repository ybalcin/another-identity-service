package account

import (
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/location"
)

func TestAddNewUser(t *testing.T) {
	users := []*user{}

	repo := new(mockRepository)
	repo.InsertNewUserFn = func(user *user) *ErrUserRepository {
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
