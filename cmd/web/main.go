package main

import (
	"forum/internal/repository/sqlite"
	"log"
)

func main() {
	// init logger

	// init db
	storage, err := sqlite.New(StoragePath)

	if err != nil {
		log.Fatal(err)
	}
	_ = storage
	// init router

	// start server
}
