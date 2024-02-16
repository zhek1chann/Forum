package app

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger

	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func New() *application {
}

func (a *application) StartSErver() {
}
