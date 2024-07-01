package main

import (
	"fmt"
	"net/http"
	"time"
)

func main () {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.yahoo.com",
		"http://www.youtube.com",
		"http://www.bing.com",
	}

	c := make(chan string)

	for _, site := range links {
		go checkLink(site, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " is down")
		c <- link
	} else {
		fmt.Println(link, " is up")
		c <- link
	}
}