package account

import (
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/location"
	"github.com/ybalcin/another-identity-service/mongo_store"
)

func TestAddNewUser(t *testing.T) {
	mongo_store.MgoConfig = mongo_store.MongoConfig{
		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
		Database: "another-identity-service-store",
	}
	mongo_store.InitMongo()

	user_repository := NewUserRepository()

	new_user := CreateNewUser(NewUserId(), "yiğitcan2", "balçın", "ybalcin", "ybalcin@", "1234567", time.Now().UTC(), "5343369694",
		true, &location.Address{
			Country: "tr",
			City:    "ist",
			County:  "küçükçekmece",
		})

	if err := user_repository.AddNewUser(new_user); err != nil {
		t.Errorf("addnewusererror: %v", err)
	}
}
