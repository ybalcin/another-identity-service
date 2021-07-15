package startup

import (
	"testing"

	"github.com/ybalcin/another-identity-service/store"
)

func TestInitConfig(t *testing.T) {
	initConfig()

	if AppConfig.Database == "" || AppConfig.MongoUri == "" || AppConfig.Port == "" {
		t.Errorf("config cannot initialized!")
	}
}

func TestInitMgoConfig(t *testing.T) {
	initMongo()

	if store.MgoConfig.Database == "" || store.MgoConfig.Uri == "" {
		t.Errorf("mgo config cannot initialized!")
	}
}
