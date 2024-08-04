package render

import (
	"MainApp/pkg/config"
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config.AppConfig

// set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// get the template cache from the app config
	tc := app.TemplateCache

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get tempalte from template cache")
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if (err != nil) {
		return myCache, err
	}

	// range through templates
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if (err != nil) {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if (err != nil) {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if (err != nil) {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}