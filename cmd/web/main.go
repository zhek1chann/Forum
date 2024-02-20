package main

import (
	"forum/internal/config"
	"forum/internal/handlers"
	"forum/internal/repository/sqlite"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// type application struct {
// 	errorLog *log.Logger
// 	infoLog  *log.Logger

// 	// templateCache  map[string]*template.Template
// 	// formDecoder    *form.Decoder
// 	// sessionManager *scs.SessionManager
// }

func main() {

	infoLog := log.New(os.Stdout, "\u001b[32mINFO\t\u001b[0m", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "\u001b[31mERROR\t\u001b[0m", log.Ldate|log.Ltime|log.Lshortfile)

	cfg := config.MustLoad()

	// init db
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Fatal(err)
	}
	_ = storage
	srv := &http.Server{
		Addr:         cfg.Address,
		ErrorLog:     errLog,
		Handler:      handlers.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on https://localhost%s", cfg.Address)
	srv.ListenAndServe()

}
