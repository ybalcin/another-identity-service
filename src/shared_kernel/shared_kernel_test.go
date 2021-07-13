package shared_kernel

import (
	"testing"

	"github.com/ybalcin/another-identity-service/mongo_store"
)

func TestInitConfig(t *testing.T) {
	initConfig()

	if AppConfig.Database == "" || AppConfig.MongoUri == "" || AppConfig.Port == "" {
		t.Errorf("config cannot initialized!")
	}
}

func TestInitMgoConfig(t *testing.T) {
	initMongo()

	if mongo_store.MgoConfig.Database == "" || mongo_store.MgoConfig.Uri == "" {
		t.Errorf("mgo config cannot initialized!")
	}
}
