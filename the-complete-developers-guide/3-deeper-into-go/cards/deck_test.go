package main

import (
	// "fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckLength := 52

	if len(d) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(d))
	}

	if d[0] != "♠J" {
		t.Errorf("Expected first card of '♠J' but got %v.", d[0])
	}

	if d[len(d)-1] != "♣10" {
		t.Errorf("Expected last card '♣10' but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDecFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	// test 1
	deckLength := 52
	if len(loadedDeck) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(loadedDeck))
	}

	os.Remove("_decktesting")
}
