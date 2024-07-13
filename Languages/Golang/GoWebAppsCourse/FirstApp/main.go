package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":3000"

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing tempalte: ", err)
		return
	}
}

func main(){
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}