package authorization

import (
	"context"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	RoleRepository interface {
		GetAll() ([]role, *common.FriendlyError)
	}
	roleRepository struct {
		roles *mongo.Collection
	}
)

func (r *roleRepository) GetAll() ([]role, *common.FriendlyError) {
	var roles []role
	ctx := context.Background()

	cursor, err := r.roles.Find(ctx, bson.M{})
	if err != nil {
		// log
		return nil, &common.FriendlyError{}
	}

	if err = cursor.All(ctx, &roles); err != nil {
		// log
		return nil, &common.FriendlyError{}
	}

	return roles, nil
}

func NewRoleRepository() RoleRepository {
	mgoStore := store.GetMgoStore()
	roleRepo := roleRepository{
		roles: mgoStore.Db.Collection(role_collection),
	}

	return &roleRepo
}
