package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Env         string `json:"env"`
	StoragePath string `json:"db_path"`
	Address     string `json:"address"`
}

func MustLoad(configPath string) *Config {
	if configPath == "" {
		log.Fatal("Config does not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatal("config file doesn't exist")
	}
	var cfg Config
	if err := ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("can't read config file")
	}

	return &cfg
}

func ReadConfig(configFilepath string, config *Config) error {
	configJsonData, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		return err
	}

	// fils the struct config which was defined above
	if err := json.Unmarshal(configJsonData, config); err != nil {
		return err
	}

	return nil
}
