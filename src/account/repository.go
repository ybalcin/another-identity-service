package account

import (
	"context"
	"log"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ERR_INSERT string = "[log_account_repository_insertnewuser_insertone]"
)

type (
	UserRepository interface {
		//	InsertNewUser inserts new user to collection
		InsertNewUser(user *user) *common.FriendlyError
		//	Adds new role to user
		AddNewRole(user *user, roleName string) *common.FriendlyError
		// Gets user by id
		GetUserById(id userId) (*user, *common.FriendlyError)
	}
	userRepository struct {
		users *mongo.Collection
	}
)

func (r *userRepository) InsertNewUser(user *user) *common.FriendlyError {

	if _, err := r.users.InsertOne(context.Background(), user); err != nil {
		log.Fatalf(ERR_INSERT+": %s", err)
		return &common.FriendlyError{
			Message:        ERR_INSERT,
			DevMessage:     ERR_INSERT,
			InnerException: err,
		}
	}

	return nil
}

func (r *userRepository) AddNewRole(user *user, roleName string) *common.FriendlyError {
	if err := user.AddRole(roleName); err != nil {
		return err
	}

	return r.UpdateRole(user, context.Background())
}

func (r *userRepository) GetUserById(id userId) (*user, *common.FriendlyError) {
	user := new(user)

	if err := r.users.FindOne(context.Background(), bson.M{"_id": primitive.ObjectID(id)}).Decode(user); err != nil {
		return nil, &common.FriendlyError{
			Message:        err.Error(),
			DevMessage:     err.Error(),
			InnerException: err,
		}
	}

	return user, nil
}

func (r *userRepository) UpdateRole(user *user, ctx context.Context) *common.FriendlyError {

	if _, err := r.users.UpdateByID(ctx, primitive.ObjectID(user.UserId), bson.D{
		{Key: "$set", Value: bson.D{{Key: "roles", Value: user.Roles}}},
	}); err != nil {
		return &common.FriendlyError{
			Message:        err.Error(),
			InnerException: err,
			DevMessage:     err.Error(),
		}
	}

	return nil
}

//	NewUserRepository gets new user repository
func NewUserRepository() UserRepository {
	mgoStore := store.GetMgoStore()
	userRepo := userRepository{
		users: mgoStore.Db.Collection(user_collection),
	}
	return &userRepo
}
