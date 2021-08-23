package authorization

import "github.com/ybalcin/another-identity-service/common"

type (
	Service interface {
		GetAll() ([]role, *common.FriendlyError)
		AddNewRole(roleId roleId, roleName string) *common.FriendlyError
	}
	service struct {
		roles RoleRepository
	}
)

func (s *service) GetAll() ([]role, *common.FriendlyError) {
	roles, err := s.roles.GetAll()
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *service) AddNewRole(roleId roleId, roleName string) *common.FriendlyError {
	role := NewRole(roleId, roleName)

	if err := s.roles.InsertNewRole(role); err != nil {
		return err
	}

	return nil
}

// NewService service initializing constructor
func NewService(roleRepository RoleRepository) Service {
	return &service{
		roles: roleRepository,
	}
}
