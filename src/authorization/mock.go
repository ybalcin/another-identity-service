package authorization

import "github.com/ybalcin/another-identity-service/common"

type (
	mockRepository struct {
		InsertNewRoleFn      func(role *role) *common.FriendlyError
		InsertNewRoleInvoked bool

		GetAllFn      func() ([]role, *common.FriendlyError)
		GetAllInvoked bool
	}
)

func (m *mockRepository) InsertNewRole(role *role) *common.FriendlyError {
	m.InsertNewRoleInvoked = true
	return m.InsertNewRoleFn(role)
}

func (m *mockRepository) GetAll() ([]role, *common.FriendlyError) {
	m.GetAllInvoked = true
	return m.GetAllFn()
}
