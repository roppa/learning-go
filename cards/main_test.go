package main

import "testing"

func TestCreateDeck(t *testing.T) {
	cards := createDeck()
	got := len(cards)
	want := 52
	if got != want {
		t.Errorf("Deck error, got %v, want %v", got, want)
	}
}

func TestCreatePlayer(t *testing.T) {
	want := "Mark"
	player := createPlayer("Mark")
	if player.name != want {
		t.Errorf("Player error, got %v, want %v", player.name, want)
	}
}
