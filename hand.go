package main

type Hand []*Card

func NewHand(cards ...*Card) Hand {
	return Hand(cards)
}
