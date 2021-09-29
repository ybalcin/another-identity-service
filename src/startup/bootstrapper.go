package startup

import (
	"github.com/ybalcin/another-identity-service/store"
)

// 	initialization sequence must not disturbed
func init() {
	//	Initialize app config
	initConfig()

	//	Initialize mongo
	initMongo()
}

//	BootstrapperInit starts application requirements
// 	initialization sequence must not disturbed
//func BootstrapperInit() {
//	//	Initialize app config
//	initConfig()
//
//	//	Initialize mongo
//	initMongo()
//}

func initMongo() {
	//	Set config
	store.MgoConfig = store.MongoConfig{
		Uri:      AppConfig.MongoUri,
		Database: AppConfig.Database,
	}
	//	Initialize mongo store
	store.InitMongo()
}
