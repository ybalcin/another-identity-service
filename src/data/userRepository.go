package data

import (
	"context"
	"log"
	"os/user"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ERR_INSERT string = "[log_userrepository_insert_insertone]"
)

type (
	ErrUserRepository struct {
		Message        string
		InnerException error
	}
)

type UserRepository interface {
	//	AddNewUser adds new user
	AddNewUser(user *user.User) *ErrUserRepository
}

type userRepository struct {
	users *mongo.Collection
}

//	Insert insert one user to collection
func (r *userRepository) AddNewUser(user *user.User) *ErrUserRepository {
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
func NewUserRepository() *userRepository {
	userRepo := userRepository{
		users: mgoStore.Db.Collection("users"),
	}
	return &userRepo
}
