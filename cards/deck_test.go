package main

import (
	"fmt"
	"testing"
)

//what do I care about the code I've written

func TestNewDeck(t *testing.T) {
	d := newDeck()

	fmt.Println(len(d))

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Sparece, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}

}
