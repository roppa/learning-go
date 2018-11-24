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

	playRound(&dealer, &player1, &deck)

	fmt.Printf("%v cards left\n", len(deck))
}

func playRound(dealer *Player, player *Player, deck *Deck) {
	dealPlayer(dealer, deck)
	dealPlayer(player, deck)
	dealPlayer(dealer, deck)
	dealPlayer(player, deck)

	dealerMin, dealerMax := dealer.getMinMax()
	playerMin, playerMax := player.getMinMax()

	fmt.Println("Dealing ...")
	fmt.Printf("Dealer min: %v, max: %v\n", dealerMin, dealerMax)
	fmt.Printf("Player min: %v, max: %v\n", playerMin, playerMax)

	for {
		if playerMin < 16 || playerMax < 16 {
			dealPlayer(player, deck)
			playerMin, playerMax = player.getMinMax()
			fmt.Printf("Player min: %v, max: %v\n\n", playerMin, playerMax)
		}

		if dealerMin < 16 || dealerMax < 16 {
			dealPlayer(dealer, deck)
			dealerMin, dealerMax = dealer.getMinMax()
			fmt.Printf("Dealer min: %v, max: %v\n", dealerMin, dealerMax)
		}

		fmt.Printf("%v cards left\n", len(*deck))

		if (playerMin >= 16 && playerMax >= 16) && (dealerMin >= 16 && dealerMax >= 16) {
			break
		}
	}

	var dealerHighest int
	var playerHighest int
	if dealerMax <= 21 {
		dealerHighest = dealerMax
	} else {
		dealerHighest = dealerMin
	}

	if playerMax <= 21 {
		playerHighest = playerMax
	} else {
		playerHighest = playerMin
	}

	if dealerHighest > 21 && playerHighest > 21 {
		fmt.Println("Both players bust")
	} else if dealerHighest > 21 && playerHighest <= 21 {
		fmt.Printf("Dealer bust: %v, player winner: %v\n", dealerHighest, playerHighest)
	} else if playerHighest > 21 && dealerHighest <= 21 {
		fmt.Printf("Player bust: %v, dealer winner: %v\n", playerHighest, dealerHighest)
	} else if dealerHighest > playerHighest || dealerHighest > 21 {
		fmt.Printf("Dealer wins: %v, player: %v\n", dealerHighest, playerHighest)
	} else if dealerHighest == playerHighest {
		fmt.Printf("No winner, dealer: %v, player: %v\n", dealerHighest, playerHighest)
	} else {
		fmt.Printf("Player wins: %v, dealer: %v\n", playerHighest, dealerHighest)
	}

	fmt.Println("Game over")
}
