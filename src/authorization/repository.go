package authorization

import (
	"context"
	"log"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ERR_INSERT string = "[log_authorization_repository_insertnewrole_insertone]"
)

type (
	RoleRepository interface {
		GetAll() ([]role, *common.FriendlyError)
		InsertNewRole(role *role) *common.FriendlyError
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

func (r *roleRepository) InsertNewRole(role *role) *common.FriendlyError {
	_, err := r.roles.InsertOne(context.Background(), role)
	if err != nil {
		log.Fatalf(ERR_INSERT+": %s", err)
		return &common.FriendlyError{
			Message:        ERR_INSERT,
			DevMessage:     ERR_INSERT,
			InnerException: err,
		}
	}

	return nil
}

func NewRoleRepository() RoleRepository {
	mgoStore := store.GetMgoStore()
	roleRepo := roleRepository{
		roles: mgoStore.Db.Collection(role_collection),
	}

	return &roleRepo
}
