package main

type Card struct {
	suit Suit
	rank Rank
}

type Suit int
type Rank int

const (
	clubs Suit = iota
	diamonds
	hearts
	spades
)

const (
	rank_Joker Rank = iota
	rank_Ace
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

var suits = []Suit{clubs, diamonds, hearts, spades}
var ranks = []Rank{rank_Ace, rank_2, rank_3, rank_4, rank_5, rank_6, rank_7, rank_8, rank_9, rank_10, rank_Jack, rank_Queen, rank_King}

func (c *Card) String() string {
	if c.rank == rank_Joker {
		return c.rank.String()
	}
	return c.rank.String() + c.suit.String()
}

func (s *Suit) String() string {
	switch *s {
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

func (r *Rank) String() string {
	switch *r {
	case rank_Joker:
		return "Jkr"
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
