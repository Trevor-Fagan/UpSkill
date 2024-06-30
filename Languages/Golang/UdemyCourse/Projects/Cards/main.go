package main

import "fmt"

func main () {
	cards := newDeck()

	cards.saveToFile("test.txt")
	newDeck := newDeckFromFile("test.txt")

	newDeck.print()
	fmt.Println("----------")
	newDeck.shuffle()
	newDeck.print()
}