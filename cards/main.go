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

// Deck - 52 cards, random order
type Deck []Card

var names = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
var suites = []string{"Hearts", "Spades", "Diamonds", "Clubs"}
var values = [][]int{{1, 11}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}, {10}, {10}, {10}}

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

// Player - will have a name and a hand
type Player struct {
	name string
	hand []Card
}

func (p *Player) getMinMax() (min, max int) {
	minTotal, maxTotal := 0, 0
	for _, card := range p.hand {
		minTotal += card.value[0]
		if len(card.value) == 1 {
			maxTotal += card.value[0]
		} else {
			maxTotal += card.value[1]
		}
	}
	return minTotal, maxTotal
}

func (p *Player) twist(card Card) {
	p.hand = append(p.hand, card)
}

func createPlayer(name string) Player {
	return Player{name: name}
}

func dealPlayer(player *Player, deck *Deck) {
	card := (*deck)[len((*deck))-1]
	// (*player).hand = append((*player).hand, card)
	player.twist(card)
	*deck = (*deck)[:len((*deck))-1]
}

func main() {
	deck := createDeck()
	dealer := createPlayer("Dealer")
	player1 := createPlayer("Mark")

	fmt.Printf("%v cards left\n", len(deck))

	dealPlayer(&dealer, &deck)
	dealPlayer(&player1, &deck)
	dealPlayer(&dealer, &deck)
	dealPlayer(&player1, &deck)

	dealerMin, dealerMax := dealer.getMinMax()
	player1Min, player1Max := player1.getMinMax()

	fmt.Printf("%v cards left\n", len(deck))
	fmt.Printf("Dealer min: %v, max: %v\n", dealerMin, dealerMax)
	fmt.Printf("Player min: %v, max: %v\n", player1Min, player1Max)

	if player1Min < 16 || player1Max < 16 {
		dealPlayer(&player1, &deck)
	}

	if dealerMin < 16 || dealerMax < 16 {
		dealPlayer(&dealer, &deck)
	}

	dealerMin, dealerMax = dealer.getMinMax()
	player1Min, player1Max = player1.getMinMax()

	fmt.Printf("%v cards left\n", len(deck))
	fmt.Printf("Dealer min: %v, max: %v\n", dealerMin, dealerMax)
	fmt.Printf("Player min: %v, max: %v\n", player1Min, player1Max)
}
