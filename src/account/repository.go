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
		//	Gets user by id
		GetUserById(id userId) (*user, *common.FriendlyError)
		//	Update one by fields
		UpdateOneByFields(user *user, fields []string) *common.FriendlyError
		//	Update fully
		UpdateOne(user *user) *common.FriendlyError
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

func (r *userRepository) UpdateOneByFields(user *user, fields []string) *common.FriendlyError {
	var updates bson.D
	for _, field := range fields {
		updates = append(updates, bson.E{Key: field, Value: user.GetFieldValue(field)})
	}

	return updateById(r.users, user, updates)
}

func (r *userRepository) UpdateOne(user *user) *common.FriendlyError {

	if _, err := r.users.UpdateByID(context.Background(), primitive.ObjectID(user.UserId), user); err != nil {
		return &common.FriendlyError{
			Message:        err.Error(),
			DevMessage:     err.Error(),
			InnerException: err,
		}
	}

	return nil
}

func updateById(userCollection *mongo.Collection, user *user, updates bson.D) *common.FriendlyError {

	if _, err := userCollection.UpdateByID(context.Background(), primitive.ObjectID(user.UserId), bson.D{
		{Key: "$set", Value: updates},
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
