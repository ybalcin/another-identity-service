package shared_kernel

import (
	"testing"

	"github.com/ybalcin/another-identity-service/data"
)

func TestInitConfig(t *testing.T) {
	initConfig()

	if AppConfig.Database == "" || AppConfig.MongoUri == "" || AppConfig.Port == "" {
		t.Errorf("config cannot initialized!")
	}
}

func TestInitMgoConfig(t *testing.T) {
	initMongo()

	if data.MgoConfig.Database == "" || data.MgoConfig.Uri == "" {
		t.Errorf("mgo config cannot initialized!")
	}
}
