package main

type Hand []Card

func NewHand(cards ...Card) Hand {
	return Hand(cards)
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

type EvalMatrix []byte

func NewMatrix() EvalMatrix {
	return EvalMatrix(make([]byte, 14))
}

func eval(hand Hand) int {
	println("len of hand:", len(hand))
	return 0
}
