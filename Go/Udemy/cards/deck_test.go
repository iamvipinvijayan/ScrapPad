package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected Deck length of 52 but got %d", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := loadDeckFromFile("_decktesting")
	loadedDeck.print()

	if len(loadedDeck) != 50 {
		t.Errorf("Expected 50 cards in deck, But got %d", len(loadedDeck))
	}

}
