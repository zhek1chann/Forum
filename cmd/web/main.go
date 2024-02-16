package main

import (
	"fmt"
	"log"
	"os"
)

type application struct {
	infoLog log.Logger
	errLog  log.Logger
}

type config struct {
	addr string
	env  string
	dsn  string
}

func main() {
	// addr := flag.String("addr", ":8080", "http network address")
	// env := flag.String("env", "dev", "dev|stage|prod")
	// dsn := flag.String("dsn", "./data/storage.db", "SQLite path")

	addr := os.Getenv("addr")
	env := os.Getenv("env")
	dsn := os.Getenv("dsn")

	fmt.Println(addr, env, dsn)

	// infoLog := log.New(os.Stdout, "\\e[0;32mINFO\t\\e[0m", log.Ldate|log.Ltime)
	// errLog := log.New(os.Stdout, "\\e[0;31m ERROR\t\\e[0m", log.Ldate|log.Ltime|log.Lshortfile)

	// type application struct{

	// }

	// init db
	// storage, err := sqlite.New(dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _ = storage
	// fmt.Println(os.Getenv("HELLO"))

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
