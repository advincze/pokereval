package main

import (
	"bytes"
)

type Hand struct {
	cards []*Card
}

func NewHand(cards ...*Card) *Hand {
	return &Hand{cards: cards}
}

func (h *Hand) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{{Cards")
	for _, card := range h.cards {
		buffer.WriteString("|")
		buffer.WriteString(card.String())
	}
	buffer.WriteString("}}")
	return buffer.String()
}
