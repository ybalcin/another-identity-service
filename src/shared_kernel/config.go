package shared_kernel

import (
	"encoding/json"
	"log"
	"os"
)

// AppConfig keeps application configs
var AppConfig config

type config struct {
	Port     string
	MongoUri string
	Database string
}

// initConfg initialize app config from json
func initConfig() {
	var err error
	var file *os.File
	file, err = os.Open("config.json")
	if err != nil {
		log.Fatalf("[log_config_initconfig_open]: %s", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	AppConfig = config{}
	if err = decoder.Decode(&AppConfig); err != nil {
		// log
		log.Fatalf("[log_config_initconfig_decode]: %s", err)
		return
	}
}
