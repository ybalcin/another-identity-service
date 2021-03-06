package account

import (
	"context"
	"log"
	"strings"

	"github.com/ybalcin/another-identity-service/common"
	"github.com/ybalcin/another-identity-service/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	err_insert string = "[log_account_repository_insertnewuser_insertone]"
)

type (
	IUserRepository interface {
		//	InsertNewUser inserts new user to collection
		InsertNewUser(user *user) *common.FriendlyError
		//	Gets user by id
		GetUserById(id userId) (*user, *common.FriendlyError)
		//	Update one by fields
		UpdateOneByFields(user *user, fields []string) *common.FriendlyError
		//	Update fully
		UpdateOne(user *user) *common.FriendlyError
		// 	GetList gets user list
		GetUserList() ([]*user, *common.FriendlyError)
	}
	userRepository struct {
		users *mongo.Collection
	}
)

func (r *userRepository) InsertNewUser(user *user) *common.FriendlyError {

	if _, err := r.users.InsertOne(context.Background(), user); err != nil {
		log.Fatalf(err_insert+": %s", err)
		return &common.FriendlyError{
			Message:        err_insert,
			DevMessage:     err_insert,
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
		updates = append(updates, bson.E{Key: strings.ToLower(field), Value: user.GetFieldValue(field)})
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

func (r *userRepository) GetUserList() ([]*user, *common.FriendlyError) {
	var users []*user

	cursor, err := r.users.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var user user
		err = cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, &user)
	}

	if err = cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
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

//	NewUserRepository user repository initializing constructor
func NewUserRepository() IUserRepository {
	mgoStore := store.GetMgoStore()
	userRepo := userRepository{
		users: mgoStore.Db.Collection(user_collection),
	}
	return &userRepo
}
