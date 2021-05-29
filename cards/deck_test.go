package main

import (
	"os"
	"testing"
)

// t is the test handler
// Testing the newDeck()
func TestNewDeck(t *testing.T) {
	deck := newDeck()

	// Checking whether the length of the deck is 16
	if len(deck) != 16 {
		t.Errorf("Expected value of 16, got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected value of Ace of Spades, got %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Expected value of Four of Clubs, got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	// Remove any test files present in the folder with name "_decktesting"
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected value of 16, got %v", len(loadedDeck))
	}

	// Remove the "_decktesting" file after completion of the test
	os.Remove("_decktesting")

}
