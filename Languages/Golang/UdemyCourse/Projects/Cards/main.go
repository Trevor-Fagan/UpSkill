package main

func main () {
	cards := newDeck()

	cards.saveToFile("test.txt")
	newDeckFromFile("test2.txt")
}