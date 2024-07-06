package main

import "fmt"

func main () {
	colors := map[string]string{
		"red": "ff0000",
		"green": "4872ds",
		"orange": "j2hsdj",
	}

	printMap(colors)

	delete(colors, "red")

	printMap(colors)
}

func printMap(myMap map[string]string) {
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}