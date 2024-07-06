package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numPool = 10

func randomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

func calculateValue(intChan chan int) {
	randomNumber := randomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go calculateValue(intChan)

	outputNum := <- intChan
	fmt.Println(outputNum)
}