package main

import (
	"bytes"
)

type Card struct {
	value int
}

func NewCard(suit Suit, rank Rank) Card {
	card := new(Card)
	card.value = int(suit)*13 + int(rank)

	return *card
}

func ParseCards(cardNames ...string) []Card {
	cards := make([]Card, len(cardNames))
	for i, cardName := range cardNames {
		cards[i] = ParseCard(cardName)
	}
	return cards
}

func ParseCard(cardName string) Card {
	suit := ParseSuit(cardName[0])
	rank := ParseRank(cardName[1:])
	return NewCard(suit, rank)
}

func ParseSuit(suitName byte) Suit {
	switch suitName {
	case 'C':
		return clubs
	case 'H':
		return hearts
	case 'D':
		return diamonds
	case 'S':
		return spades
	}
	panic("no other suits")
}

func ParseRank(rankName string) Rank {
	switch rankName {
	case "A":
		return rank_Ace
	case "2":
		return rank_2
	case "3":
		return rank_3
	case "4":
		return rank_4
	case "5":
		return rank_5
	case "6":
		return rank_6
	case "7":
		return rank_7
	case "8":
		return rank_8
	case "9":
		return rank_9
	case "10":
		return rank_10
	case "J":
		return rank_Jack
	case "Q":
		return rank_Queen
	case "K":
		return rank_King
	}
	panic("no other ranks")
}

func (c Card) getSuit() Suit {
	return Suit(int(c.value) / 13)
}

func (c Card) getRank() Rank {
	return Rank(int(c.value) % 13)
}

func (c *Card) String() string {
	return c.getRank().String() + c.getSuit().String()
}

type Suit int

const (
	clubs Suit = iota
	diamonds
	hearts
	spades
)

var suits = []Suit{clubs, diamonds, hearts, spades}

func (s Suit) String() string {
	switch s {
	case clubs:
		return "♣"
	case diamonds:
		return "♦"
	case hearts:
		return "♥"
	case spades:
		return "♠"
	}
	panic("no other suits")
}

type Rank int

const (
	rank_Ace Rank = iota
	rank_2
	rank_3
	rank_4
	rank_5
	rank_6
	rank_7
	rank_8
	rank_9
	rank_10
	rank_Jack
	rank_Queen
	rank_King
	rank_AceHigh
)

var ranks = []Rank{rank_Ace, rank_2, rank_3, rank_4, rank_5, rank_6, rank_7, rank_8, rank_9, rank_10, rank_Jack, rank_Queen, rank_King}

func (r Rank) String() string {
	switch r {
	case rank_Ace:
		return "A"
	case rank_2:
		return "2"
	case rank_3:
		return "3"
	case rank_4:
		return "4"
	case rank_5:
		return "5"
	case rank_6:
		return "6"
	case rank_7:
		return "7"
	case rank_8:
		return "8"
	case rank_9:
		return "9"
	case rank_10:
		return "10"
	case rank_Jack:
		return "J"
	case rank_Queen:
		return "Q"
	case rank_King:
		return "K"
	case rank_AceHigh:
		return "A2"
	}
	panic("no other ranks")
}

type Cards []Card

func (cards Cards) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("{{Cards")
	for _, card := range cards {
		buffer.WriteString("|")
		buffer.WriteString(card.String())
	}
	buffer.WriteString("}}")
	return buffer.String()
}
