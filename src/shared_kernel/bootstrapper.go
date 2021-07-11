package shared_kernel

import (
	"github.com/ybalcin/another-identity-service/data"
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
	data.MgoConfig = data.MongoConfig{
		Uri:      AppConfig.MongoUri,
		Database: AppConfig.Database,
	}
	//	Initialize mongo store
	data.InitMongo()
}
