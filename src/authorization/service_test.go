package authorization

import (
	"testing"

	"github.com/ybalcin/another-identity-service/common"
)

func TestService_AddNewRole(t *testing.T) {
	var roles []*role

	repo := new(mockRepository)
	repo.InsertNewRoleFn = func(role *role) *common.FriendlyError {
		roles = append(roles, role)
		return nil
	}

	service := NewService(repo)
	if err := service.AddNewRole(NewRoleId(), "guest"); err != nil {
		t.FailNow()
	}

	isExist := false
	for _, role := range roles {
		if role.Name == "guest" {
			isExist = true
		}
	}

	if !isExist {
		t.FailNow()
	}
}
