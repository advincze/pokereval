package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(int64(8))
	deck := NewDeck()
	fmt.Printf("deck: %v\n", Cards(deck))
	deck.Shuffle()
	fmt.Printf("deck: %v\n", Cards(deck))
	deck, cards := deck.giveTopCards(7)
	fmt.Printf("hand: %v\n", Cards(cards))
	eval(Hand(cards))
}
