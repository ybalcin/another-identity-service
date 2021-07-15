package account

import (
	"context"
	"log"

	"github.com/ybalcin/another-identity-service/store"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ERR_INSERT string = "[log_account_repository_insert_insertone]"
)

type (
	ErrUserRepository struct {
		Message        string
		InnerException error
	}
)

type UserRepository interface {
	//	InsertNewUser inserts new user to collection
	InsertNewUser(user *user) *ErrUserRepository
}

type userRepository struct {
	users *mongo.Collection
}

func (r *userRepository) InsertNewUser(user *user) *ErrUserRepository {
	_, err := r.users.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatalf(ERR_INSERT+": %s", err)
		return &ErrUserRepository{
			Message:        ERR_INSERT,
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
