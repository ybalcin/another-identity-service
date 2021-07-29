package account

import "github.com/ybalcin/another-identity-service/common"

type (
	mockRepository struct {
		InsertNewUserFn      func(user *user) *common.FriendlyError
		InsertNewUserInvoked bool
	}
)

func (r *mockRepository) InsertNewUser(user *user) *common.FriendlyError {
	r.InsertNewUserInvoked = true
	return r.InsertNewUserFn(user)
}
