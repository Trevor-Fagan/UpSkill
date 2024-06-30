package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Test length of deck
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	// Test first element in deck
	if (d[0] != "Ace of Spades") {
		t.Errorf("Expected Ace of Spades as first element, got %s", d[0])
	}

	// Test last element in deck
	if (d[15] != "Four of Clubs") {
		t.Errorf("Expected Four of Clubs as last element, got %s", d[len(d) - 1])
	}
}

func TestLoadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")
	ld := newDeckFromFile("_decktesting")
	
	if len(ld) != 16 {
		t.Errorf("Expected deck length of 16, got %d", len(ld))
	}

	os.Remove("_decktesting")
}