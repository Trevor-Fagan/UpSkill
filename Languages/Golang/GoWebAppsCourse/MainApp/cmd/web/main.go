package main

import (
	"MainApp/pkg/config"
	"MainApp/pkg/handlers"
	"MainApp/pkg/render"
	"fmt"
	"net/http"
	"log"
)

const portNumber = ":3000"

func main(){
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}