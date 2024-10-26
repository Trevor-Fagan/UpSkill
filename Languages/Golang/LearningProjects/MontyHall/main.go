package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
Requirements:
- For powers of 10, run test for Monty Hall problem
- Output the percentage of times you get it right by switching vs staying
*/
func main() {
	for i := 0; i < 5; i++ {
		numSimulations := math.Pow(10, float64(i))
		doors := 3
		correctGuesses := 0

		for j := 0; j < int(numSimulations); j++ {
			computerGuess := 1 + rand.Intn(doors)
			userGuess := 1 + rand.Intn(doors)

			if userGuess == computerGuess {
				correctGuesses++
			}
		}

		fmt.Printf("%d: %f\n", int(numSimulations), 100*float64(correctGuesses)/numSimulations)
	}
}
