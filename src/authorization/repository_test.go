package authorization

import (
	"testing"

	"github.com/ybalcin/another-identity-service/store"
)

func TestGetAll(t *testing.T) {
	store.MgoConfig = store.MongoConfig{
		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
		Database: "another-identity-service-store",
	}
	store.InitMongo()

	repository := NewRoleRepository()

	_, err := repository.GetAll()
	if err != nil {
		t.FailNow()
	}
}

func TestInsertNewRole(t *testing.T) {
	store.MgoConfig = store.MongoConfig{
		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
		Database: "another-identity-service-store",
	}
	store.InitMongo()

	repository := NewRoleRepository()

	role := CreateNewRole(NewRoleId(), "guest")
	if err := repository.InsertNewRole(role); err != nil {
		t.Error(err)
	}
}
