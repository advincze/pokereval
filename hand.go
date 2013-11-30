package main

import (
	"bytes"
)

type Hand []Card

func NewHand(cards ...Card) Hand {
	return Hand(cards)
}

func (h Hand) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{{Cards")
	for _, card := range h {
		buffer.WriteString("|")
		buffer.WriteString(card.String())
	}
	buffer.WriteString("}}")
	return buffer.String()
}

type Combination int

const (
	high_card Combination = iota
	pair
	two_pairs
	three_of_a_kind
	straight
	flush
	full_house
	four_of_a_kind
	straight_flush
)
