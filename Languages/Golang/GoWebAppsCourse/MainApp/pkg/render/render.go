package render

import (
	"text/template"
	"fmt"
	"net/http"
	"os"
)

// Render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("templates/" + tmpl, "templates/base.layout.tmpl")
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