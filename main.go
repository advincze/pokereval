package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := NewDeck()
	rand.Seed(time.Now().UTC().UnixNano())
	deck.Shuffle()
	fmt.Printf("%v\n", deck)
	hand := deck.giveRoundRobinHands(1)[0]
	fmt.Printf("%v\n", hand)
}
