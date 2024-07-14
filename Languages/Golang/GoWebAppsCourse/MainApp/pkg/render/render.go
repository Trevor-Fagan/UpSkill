package render

import (
	"text/template"
	"fmt"
	"net/http"
)

// Render HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("../../templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing tempalte: ", err)
		return
	}
}