package main

import (
	"bytes"
	"math/rand"
)

type Deck []Card

func NewDeck() Deck {
	deck := Deck(make([]Card, 0, 52))
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, NewCard(suit, rank))
		}
	}
	return deck
}

func (d Deck) giveRoundRobinHands(numberOfHands int) []Hand {
	hands := make([]Hand, numberOfHands)
	for i := 0; i < numberOfHands; i++ {
		hands[i] = NewHand(d.giveTopCard())
	}
	for i := 0; i < numberOfHands; i++ {
		hands[i] = append(hands[i], d.giveTopCard())
	}
	return hands
}

func (d Deck) giveTopCard() Card {
	return d.giveCard(0)
}

func (d Deck) giveCard(index int) Card {
	card := d[index]
	d = append(d[:index], d[index+1:]...)
	return card
}

func (d Deck) Shuffle() {
	for i := len(d) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			d[i], d[j] = d[j], d[i]
		}
	}

}

func (d Deck) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{{Cards")
	for _, card := range d {
		buffer.WriteString("|")
		buffer.WriteString(card.String())
	}
	buffer.WriteString("}}")
	return buffer.String()
}
