package account

import (
	"github.com/ybalcin/another-identity-service/common"
)

type (
	mockRepository struct {
		InsertNewUserFn      func(user *user) *common.FriendlyError
		InsertNewUserInvoked bool

		AddNewRoleFn      func(user *user, roleName string) *common.FriendlyError
		AddNewRoleInvoked bool

		GetUserByIdFn      func(id userId) (*user, *common.FriendlyError)
		GetUserByIdInvoked bool
	}
)

func (r *mockRepository) InsertNewUser(user *user) *common.FriendlyError {
	r.InsertNewUserInvoked = true
	return r.InsertNewUserFn(user)
}

func (r *mockRepository) AddNewRole(user *user, roleName string) *common.FriendlyError {
	r.AddNewRoleInvoked = true
	return r.AddNewRoleFn(user, roleName)
}

func (r *mockRepository) GetUserById(id userId) (*user, *common.FriendlyError) {
	r.GetUserByIdInvoked = true
	return r.GetUserByIdFn(id)
}
