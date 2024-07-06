package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main() {
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)
}