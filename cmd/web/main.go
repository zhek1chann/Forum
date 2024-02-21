package main

import (
	"forum/app"
	"forum/internal/config"
	"forum/internal/handlers"
	repo "forum/internal/repository"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	infoLog := log.New(os.Stdout, "\u001b[32mINFO\t\u001b[0m", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "\u001b[31mERROR\t\u001b[0m", log.Ldate|log.Ltime|log.Lshortfile)

	cfg := config.MustLoad()

	tc, err := app.NewTemplateCache()

	if err != nil {
		errLog.Fatal(err)
	}

	app := app.New(infoLog, errLog, tc)

	_ = app
	// init db
	r, err := repo.New(cfg.StoragePath)
	if err != nil {
		log.Fatal(err)
	}
	_ = r

	srv := &http.Server{
		Addr:         cfg.Address,
		ErrorLog:     errLog,
		Handler:      handlers.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on http://localhost%s", cfg.Address)
	srv.ListenAndServe()

}
