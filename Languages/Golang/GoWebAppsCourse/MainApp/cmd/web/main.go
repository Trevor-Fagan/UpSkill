package main

import (
	"fmt"
	"net/http"
	"MainApp/pkg/handlers"
)

const portNumber = ":3000"

func main(){
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}