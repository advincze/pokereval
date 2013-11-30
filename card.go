package main

type Card int

func NewCard(suit Suit, rank Rank) Card {
	return Card(int(suit)*len(ranks) + int(rank))
}

func (c Card) getSuit() Suit {
	return Suit(int(c) / len(ranks))
}

func (c Card) getRank() Rank {
	return Rank(int(c) % len(ranks))
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
	}
	panic("no other ranks")
}
