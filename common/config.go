package common

import (
	"encoding/json"
	"log"
	"os"
)

var Config Configuration

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string

	MySQLURL string
}

func LoadConfig(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
