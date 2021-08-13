package account

import (
	"testing"
	"time"

	"github.com/ybalcin/another-identity-service/location"
	"github.com/ybalcin/another-identity-service/store"
)

func TestInsertNewUser(t *testing.T) {
	store.MgoConfig = store.MongoConfig{
		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
		Database: "another-identity-service-store",
	}
	store.InitMongo()

	user_repository := NewUserRepository()

	new_user, _ := CreateNewUser(NewUserId(), "yiğitcan2", "balçın", "ybalcin", "ybalcin@", "1234567", time.Now().UTC(), "5343369694",
		true, &location.Address{
			Country: "tr",
			City:    "ist",
			County:  "küçükçekmece",
		})

	if err := user_repository.InsertNewUser(new_user); err != nil {
		t.Errorf("addnewusererror: %v", err)
	}
}
