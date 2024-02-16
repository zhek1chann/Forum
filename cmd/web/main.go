package main

import (
	"fmt"
	"forum/internal/config"
	"forum/internal/repository/sqlite"
	"log"
)

func main() {
	// init config
	cfg := config.MustLoad("././config/config.json")

	fmt.Println(cfg)
	// init logger

	// init db
	storage, err := sqlite.New(cfg.StoragePath)

	if err != nil {
		log.Fatal(err)
	}
	_ = storage

	// init router

	// start server
}
