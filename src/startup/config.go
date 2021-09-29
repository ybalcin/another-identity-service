package startup

import (
	"encoding/json"
	"log"
	"os"
)

const (
	ERR_OPEN   string = "[log_config_initconfig_open]: %s"
	ERR_DECODE string = "[log_config_initconfig_decode]: %s"
)

// AppConfig keeps application configs
var AppConfig config

type config struct {
	Port     string
	MongoUri string
	Database string
}

// initConfig initialize app config from json
func initConfig() {
	var err error
	var file *os.File
	file, err = os.Open("config.json")
	if err != nil {
		log.Fatalf(ERR_OPEN, err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	AppConfig = config{}
	if err = decoder.Decode(&AppConfig); err != nil {
		// log
		log.Fatalf(ERR_DECODE, err)
		return
	}
}
