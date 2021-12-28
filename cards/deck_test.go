package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected Deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card as Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Jack of Diamonds" {
		t.Errorf("Expected Jack of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_deckTesting"
	err := os.Remove(filename)
	d := newDeck()
	err = d.saveToFile(filename)
	if err != nil {
		t.Errorf("Error while Saving %v", err)
	}
	loadedDeck := newDeckFromFile(filename)
	if len(loadedDeck) != 52 {
		t.Errorf("Expecting 52 cards in deck, but got %v", len(loadedDeck))
	}
	err = os.Remove(filename)
}
