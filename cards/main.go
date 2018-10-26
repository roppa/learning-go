package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Card - Playing card
type Card struct {
	name  string
	suit  string
	value []int
}

// Player - will have a name and a hand
type Player struct {
	name string
	hand []Card
}

// Deck - 52 cards, random order
type Deck []Card

var names = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var suites = []string{"Hearts", "Spades", "Diamonds", "Clubs"}
var values = [][]int{{1, 11}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}, {10}, {10}, {10}}

func main() {
	deck := createDeck()
	dealer := createPlayer("Dealer")
	player1 := createPlayer("Mark")

	deal(&dealer, &deck)
	deal(&player1, &deck)

	deal(&dealer, &deck)
	deal(&player1, &deck)

	dealerMin, dealerMax := getPlayerTotal(dealer)
	player1Min, player1Max := getPlayerTotal(player1)
	fmt.Printf("%v cards left\n", len(deck))
	fmt.Printf("Dealer min: %v, max: %v\n", dealerMin, dealerMax)
	fmt.Printf("Player min: %v, max: %v\n", player1Min, player1Max)
}

func createDeck() Deck {
	cards := make(Deck, 0)
	for _, suit := range suites {
		for i, value := range values {
			cards = append(cards, Card{
				name:  names[i],
				suit:  suit,
				value: value,
			})
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func createPlayer(name string) Player {
	return Player{name: name}
}

func deal(player *Player, deck *Deck) {
	card := (*deck)[len((*deck))-1]
	(*player).hand = append((*player).hand, card)
	*deck = (*deck)[:len((*deck))-1]
}

func getPlayerTotal(player Player) (int, int) {
	minTotal, maxTotal := 0, 0
	for _, card := range player.hand {
		minTotal += card.value[0]
		if len(card.value) == 1 {
			maxTotal += card.value[0]
		} else {
			maxTotal += card.value[1]
		}
	}
	return minTotal, maxTotal
}
