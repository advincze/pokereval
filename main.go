package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	deck := NewDeck()
	deck.Shuffle()
	fmt.Printf("%v\n", deck)
	hand := deck.giveRoundRobinHands(1)[0]
	fmt.Printf("%v\n", hand)
}

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

func (d *Deck) giveRandomHand() *Hand {
	return NewHand(d.giveRandomCard(), d.giveRandomCard())
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

func (d *Deck) giveRandomCard() *Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return d.giveCard(r.Intn(len(d.cards)))
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

type Card struct {
	suit Suit
	rank Rank
}

func (c *Card) String() string {
	return c.rank.String() + c.suit.String()
}

type Suit int
type Rank int

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
	case rank_ace:
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

const (
	clubs Suit = iota
	diamonds
	hearts
	spades
)

const (
	rank_ace Rank = iota
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
var ranks = []Rank{rank_ace, rank_2, rank_3, rank_4, rank_5, rank_6, rank_7, rank_8, rank_9, rank_10, rank_Jack, rank_Queen, rank_King}
