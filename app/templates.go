package app

import (
	"bytes"
	"fmt"
	"forum/models"
	"forum/ui"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"time"
)

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}

func (app *Application) render(w http.ResponseWriter, status int, page string, data *models.TemplateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}
	buf := new(bytes.Buffer)
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (app *Application) newTemplateData(r *http.Request) *models.TemplateData {
	return &models.TemplateData{
		CurrentYear: time.Now().Year(),
		//Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		//IsAuthenticated: app.isAuthenticated(r),
		//CSRFToken:       nosurf.Token(r),
	}
}
