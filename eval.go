package main

const NUM_CARDS = 5
const FLAG_STRAIGHT_FLUSH = 1 << 31
const FLAG_FOUR_OF_A_KIND = 1 << 30
const FLAG_FULL_HOUSE = 1 << 29
const FLAG_FLUSH = 1 << 28
const FLAG_STRAIGHT = 1 << 27
const FLAG_THREE_OF_A_KIND = 1 << 26
const FLAG_TWO_PAIRS = 1 << 25
const FLAG_ONE_PAIR = 1 << 24

type EvalResult struct {
	straight bool
	flush    bool
	four     bool
	three    bool
	pairs    int
	score    uint32
}

func eval5(hand Hand) *EvalResult {

	//precalculations
	var numInRank [NUM_RANKS]int
	var numInSuit [NUM_SUITS]int
	for _, card := range hand {
		numInRank[card.getRank()]++
		numInSuit[card.getSuit()]++
	}

	// check for flush
	flush := checkFlush(&numInSuit)

	// check for straight
	straight := checkStraight(&numInRank)

	//check four of a kind, three of a kind and pairs
	var four bool
	var three bool
	var highRank Rank
	var lowerRank Rank
	var pairs int
	var rankScore uint32
	var score uint32
	for rank := rank_2; rank <= rank_Ace; rank += 1 {

		if numInRank[rank] != 0 {
			rankScore |= 1 << rank
		}

		switch numInRank[rank] {
		case 2:
			pairs += 1
			highRank, lowerRank = lowerRank, rank+1
		case 3:
			three = true
			highRank = rank + 1
		case 4:
			four = true
			highRank = rank + 1
		}
	}

	score |= rankScore

	if straight && flush {
		score |= FLAG_STRAIGHT_FLUSH
	} else if four {
		score |= FLAG_FOUR_OF_A_KIND
	} else if three && pairs == 1 {
		score |= FLAG_FULL_HOUSE
	} else if flush {
		score |= FLAG_FLUSH
	} else if straight {
		score |= FLAG_STRAIGHT
	} else if three {
		score |= FLAG_THREE_OF_A_KIND
	}

	score |= uint32(highRank+1) << 20
	score |= uint32(lowerRank+1) << 16

	return &EvalResult{
		straight: straight,
		flush:    flush,
		four:     four,
		three:    three,
		pairs:    pairs,
		score:    score,
	}
}

func checkFlush(numInSuit *[NUM_SUITS]int) bool {
	for _, suit := range suits {
		if numInSuit[suit] == NUM_CARDS {
			return true
		}
	}
	return false
}

func checkStraight(numInRank *[NUM_RANKS]int) bool {

	var consecutive_ranks int
	rank := rank_2
	for ; numInRank[rank] == 0; rank++ {
	}
	for ; rank <= rank_Ace && numInRank[rank] != 0; rank++ {
		consecutive_ranks++
	}
	if consecutive_ranks == NUM_CARDS {
		return true

	} else if consecutive_ranks == NUM_CARDS-1 && numInRank[rank_Ace] != 0 && numInRank[rank_2] != 0 && numInRank[rank_5] != 0 {
		return true
	}
	return false
}
