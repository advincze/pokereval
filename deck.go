package main

import (
	"bytes"
	"math/rand"
)

type Deck struct {
	cards []*Card
}

func NewDeck() *Deck {
	deck := &Deck{}
	deck.cards = make([]*Card, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.putCard(&Card{suit: suit, rank: rank})
		}
	}
	return deck
}

func (d *Deck) putCard(card *Card) {
	d.cards = append(d.cards, card)
}

func (d *Deck) giveRoundRobinHands(numberOfHands int) []*Hand {
	hands := make([]*Hand, numberOfHands)
	for i := 0; i < numberOfHands; i++ {
		hands[i] = NewHand(d.giveTopCard())
	}
	for i := 0; i < numberOfHands; i++ {
		hands[i].cards = append(hands[i].cards, d.giveTopCard())
	}
	return hands
}

func (d *Deck) giveTopCard() *Card {
	if len(d.cards) == 0 {
		return nil
	}
	return d.giveCard(0)
}

func (d *Deck) giveCard(index int) *Card {
	card := d.cards[index]
	d.cards = append(d.cards[:index], d.cards[index+1:]...)
	return card
}

func (d *Deck) Shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
		}
	}

}

func (d *Deck) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{{Cards")
	for _, card := range d.cards {
		buffer.WriteString("|")
		buffer.WriteString(card.String())
	}
	buffer.WriteString("}}")
	return buffer.String()
}
