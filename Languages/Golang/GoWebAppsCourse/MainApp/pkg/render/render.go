package render

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

// Render HTML templates
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("templates/"+tmpl, "templates/base.layout.tmpl")
	if err != nil {
		fmt.Println("error caught: ", err)
		os.Exit(1)
	}

	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing tempalte: ", err)
		return
	}
}

// template cache variable
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template
	_, inMap := tc[t]
	if !inMap {
		fmt.Println("creating template and adding to cache")

		// create the template
		err = createTemplateCache(t)

		if err != nil {
			fmt.Println(err)
		}
	} else {
		// return cached template
		fmt.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl
	return nil
}