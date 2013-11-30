package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := NewDeck()
	fmt.Printf("%v\n", deck)
	rand.Seed(time.Now().UTC().UnixNano())
	deck.Shuffle()
	fmt.Printf("%v\n", deck)
	hand := deck.giveRoundRobinHands(1)
	fmt.Printf("%v\n", hand)
}
