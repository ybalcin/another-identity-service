package account

import (
	"github.com/ybalcin/another-identity-service/common"
)

type (
	mockRepository struct {
		InsertNewUserFn      func(user *user) *common.FriendlyError
		InsertNewUserInvoked bool

		GetUserByIdFn      func(id userId) (*user, *common.FriendlyError)
		GetUserByIdInvoked bool

		UpdateOneByFieldsFn      func(user *user, fields []string) *common.FriendlyError
		UpdateOneByFieldsInvoked bool

		UpdateOneFn      func(user *user) *common.FriendlyError
		UpdateOneInvoked bool

		GetUserListFn      func() ([]*user, *common.FriendlyError)
		GetUserListInvoked bool
	}
)

func (r *mockRepository) InsertNewUser(user *user) *common.FriendlyError {
	r.InsertNewUserInvoked = true
	return r.InsertNewUserFn(user)
}

func (r *mockRepository) GetUserById(id userId) (*user, *common.FriendlyError) {
	r.GetUserByIdInvoked = true
	return r.GetUserByIdFn(id)
}

func (r *mockRepository) UpdateOneByFields(user *user, fields []string) *common.FriendlyError {
	r.UpdateOneByFieldsInvoked = true
	return r.UpdateOneByFieldsFn(user, fields)
}

func (r *mockRepository) UpdateOne(user *user) *common.FriendlyError {
	r.UpdateOneInvoked = true
	return r.UpdateOneFn(user)
}

func (r *mockRepository) GetUserList() ([]*user, *common.FriendlyError) {
	r.GetUserListInvoked = true
	return r.GetUserListFn()
}
