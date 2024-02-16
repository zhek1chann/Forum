package main

import (
	"forum/internal/config"
	"forum/internal/repository/sqlite"
	"log"
)

func main() {

	// infoLog := log.New(os.Stdout, "\\e[0;32mINFO\t\\e[0m", log.Ldate|log.Ltime)
	// errLog := log.New(os.Stdout, "\\e[0;31m ERROR\t\\e[0m", log.Ldate|log.Ltime|log.Lshortfile)

	cfg := config.MustLoad()

	// init db
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Fatal(err)
	}
	_ = storage
	// a := app.New()

	// a.StartSErver()

	// init router

	// start server
}

// type application struct {
// 	errorLog *log.Logger
// 	infoLog  *log.Logger

// 	templateCache  map[string]*template.Template
// 	formDecoder    *form.Decoder
// 	sessionManager *scs.SessionManager
// }
