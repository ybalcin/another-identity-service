package account

import (
	"fmt"
	"testing"

	"github.com/ybalcin/another-identity-service/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func TestInsertNewUser(t *testing.T) {
// 	store.MgoConfig = store.MongoConfig{
// 		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
// 		Database: "another-identity-accountService-store",
// 	}
// 	store.InitMongo()

// 	user_repository := NewUserRepository()

// 	new_user, _ := CreateNewUser(NewUserId(), "yiğitcan2", "balçın", "ybalcin", "ybalcin@", "1234567", time.Now().UTC(), "5343369694",
// 		true, &location.Address{
// 			Country: "tr",
// 			City:    "ist",
// 			County:  "küçükçekmece",
// 		})

// 	if err := user_repository.InsertNewUser(new_user); err != nil {
// 		t.Errorf("addnewusererror: %v", err)
// 	}
// }

func TestUserRepository_GetUserById(t *testing.T) {
	store.MgoConfig = store.MongoConfig{
		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
		Database: "another-identity-accountService-store",
	}
	store.InitMongo()

	userRepository := NewUserRepository()

	id, _ := primitive.ObjectIDFromHex("60f70c937c6ebf50d5464366")

	user, err := userRepository.GetUserById(userId(id))
	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.FailNow()
	}

	fmt.Printf("\n" + user.Username + "\n")
}

func TestUserRepository_GetUserList(t *testing.T) {
	userRepository := NewUserRepository()

	_, err := userRepository.GetUserList()
	if err != nil {
		t.Fatal(err)
	}
}
