package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected a deck length of 52, but got %d", len(d))
	}

}

func TestFirstCard(t *testing.T){
	d := newDeck()
	if d[0] != "ace of spades"{
		t.Errorf("expected the ace of spades, but got %v", d[0])
	}
}

func TestLastCard(t *testing.T){
	d := newDeck()
	if d[len(d)-1] != "king of clubssssssss"{
		t.Errorf("expected the king of clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) <= 1 {
		t.Errorf("Expected a new deck of 100 cards, only found %d cards", len(loadedDeck))
	}

	os.Remove("_decktesting")
}