package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Spider struct {
	Legs int `json:"legs"`
	Name string `json:"name"`
	Color string `json:"color"`
}

func marshallJSON(s Spider) string {
	encodedJson, err := json.Marshal(s)

	if err != nil {
		log.Println("Error unmarshalling JSON")
		os.Exit(1)
	}

	return string(encodedJson)
}

func main () {
	theJson := `
[{"legs": 8, "name": "daddy long legs", "color": "red"},
 {"legs": 6, "name": "bad guy", "color": "orange"}]
`

	var unmarshalled []Spider
	err := json.Unmarshal([]byte(theJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling JSON: ", err)
	} else {
		fmt.Println(unmarshalled)
	}

	daddyLongLegs := Spider{10, "daddy long legs", "gray"}
	encodedJson, err := json.Marshal(daddyLongLegs)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(encodedJson))
	}
}