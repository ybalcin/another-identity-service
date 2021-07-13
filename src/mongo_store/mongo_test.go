package mongo_store

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestInitMongo(t *testing.T) {
	MgoConfig = MongoConfig{
		Uri:      "mongodb+srv://identityServiceUser:HMPQ4jXPCrxEDB58@cluster0.l1pmb.mongodb.net",
		Database: "another-identity-service-store",
	}

	InitMongo()

	if err := mgoStore.Session.Ping(context.Background(), readpref.Primary()); err != nil {
		t.Errorf("mongo ping error: %s", err)
	}
}
