package config

import (
	"log"
	"os"
)

type Config struct {
	Env         string `json:"env"`
	StoragePath string `json:"db_path"`
	Address     string `json:"address"`
}

func MustLoad() *Config {
	var cfg Config
	cfg.Address = os.Getenv("addr")
	cfg.Env = os.Getenv("env")
	cfg.StoragePath = os.Getenv("storage_path")

	if cfg.Address == "" {
		log.Fatal("Addres is not set")
	}

	if cfg.Env == "" {
		log.Fatal("Envirment is not set")
	}

	if cfg.StoragePath == "" {
		log.Fatal("storage path is not set")
	}

	return &cfg
}
