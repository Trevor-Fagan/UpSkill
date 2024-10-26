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
		//correctGuessesStay := userStaySimulation(int(numSimulations))
		correctGuessesSwitch := userSwtichSimulation(int(numSimulations))

		fmt.Printf("%d: %f\n", int(numSimulations), 100*float64(correctGuessesSwitch)/numSimulations)
	}
}

func userStaySimulation(numSimulations int) int {
	doors := 3
	correctGuesses := 0

	for i := 0; i < numSimulations; i++ {
		computerGuess := 1 + rand.Intn(doors)
		userGuess := 1 + rand.Intn(doors)

		if userGuess == computerGuess {
			correctGuesses++
		}
	}

	return correctGuesses
}

func userSwtichSimulation(numSimulations int) int {
	correctGuesses := 0

	for i := 0; i < numSimulations; i++ {
		computerGuess := 1 + rand.Intn(3)
		userGuess := 1 + rand.Intn(3)

		if computerGuess != userGuess {
			genNumber := rand.Intn(2)

			if genNumber == 0 {
				correctGuesses++
			}
		}
	}

	return correctGuesses
}
