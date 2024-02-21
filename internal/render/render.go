package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var funcstions = template.FuncMap{}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./web/templates/page/*.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(funcstions).ParseFiles(page)
		if err != nil {
			fmt.Println(1)
			return myCache, err
		}

		layouts, err := filepath.Glob("./web/templates/*.html")
		if err != nil {
			fmt.Println(2)

			return myCache, err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./web/templates/*html")
			if err != nil {
				fmt.Println(3)
				return myCache, err
			}
		}

		partials, err := filepath.Glob("./web/templates/partial/*.html")
		if err != nil {
			return myCache, err
		}
		if len(partials) > 0 {
			ts, err = ts.ParseGlob("./web/templates/partial/*.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, err
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("coult not create cache")
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("coult not get template from cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}
