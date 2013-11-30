package main

import (
	"math/rand"
)

type Deck []*Card

func NewDeck() Deck {
	deck := Deck(make([]*Card, 0, 52))
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, NewCard(suit, rank))
		}
	}
	return deck
}

func (d Deck) giveRoundRobinHands(noHands, noCards int) []Hand {
	hands := make([]Hand, 0, noHands)
	var card *Card
	for j := 0; j < noCards; j++ {
		for i := 0; i < noHands; i++ {
			d, card = d.giveTopCard()
			hands[i] = append(hands[i], card)
		}
	}
	return hands
}

func (d Deck) giveTopCard() (Deck, *Card) {
	return d[1:], d[0]
}

func (d Deck) giveTopCards(number int) (Deck, []*Card) {
	return d[number:], d[:number]
}

func (d Deck) giveCard(index int) (Deck, *Card) {
	card := d[index]
	d = append(d[:index], d[index+1:]...)
	return d, card
}

func (d Deck) Shuffle() {
	for i := len(d) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			d[i], d[j] = d[j], d[i]
		}
	}
}
