package shared_kernel

import (
	"github.com/ybalcin/another-identity-service/mongo_store"
)

//	BootstrapperInit starts application requirements
// 	initialization sequence must not disturbed
func BootstrapperInit() {
	//	Initialize app config
	initConfig()

	//	Initialize mongo
	initMongo()
}

func initMongo() {
	//	Set config
	mongo_store.MgoConfig = mongo_store.MongoConfig{
		Uri:      AppConfig.MongoUri,
		Database: AppConfig.Database,
	}
	//	Initialize mongo store
	mongo_store.InitMongo()
}
