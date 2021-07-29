package account

import (
	"context"
	"log"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/store"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ERR_INSERT string = "[log_account_repository_insertnewuser_insertone]"
)

type UserRepository interface {
	//	InsertNewUser inserts new user to collection
	InsertNewUser(user *user) *common.FriendlyError
}

type userRepository struct {
	users *mongo.Collection
}

func (r *userRepository) InsertNewUser(user *user) *common.FriendlyError {

	_, err := r.users.InsertOne(context.Background(), user)
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

//	NewUserRepository gets new user repository
func NewUserRepository() UserRepository {
	mgoStore := store.GetMgoStore()
	userRepo := userRepository{
		users: mgoStore.Db.Collection(user_collection),
	}
	return &userRepo
}
