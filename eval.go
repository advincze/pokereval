package main

const NUM_CARDS = 5

type EvalResult struct {
	straight bool
	flush    bool
	four     bool
	three    bool
	pairs    int
	score    uint32
}

func eval5(hand Hand) *EvalResult {
	// println("len of hand:", len(hand))
	score := uint32(0)

	//precalculations
	var numInRank [NUM_RANKS]int
	var numInSuit [NUM_SUITS]int
	for _, card := range hand {
		numInRank[card.getRank()]++
		numInSuit[card.getSuit()]++
	}

	// check for flush
	var flush bool
	for _, suit := range suits {
		if numInSuit[suit] == NUM_CARDS {
			flush = true
			break
		}
	}

	// check for straight
	straight := checkStraight(&numInRank)

	//check four of a kind, three of a kind and pairs
	var four bool
	var three bool
	var highRank Rank
	var lowerRank Rank
	var pairs int
	var rankScore uint32
	for rank := rank_2; rank <= rank_Ace; rank += 1 {

		if numInRank[rank] != 0 {
			rankScore |= 1 << rank
		}

		switch numInRank[rank] {
		case 2:
			pairs += 1
			highRank, lowerRank = rank, highRank
		case 3:
			three = true
			highRank = rank
		case 4:
			four = true
			highRank = rank
		}
	}

	score |= rankScore

	if straight && flush {
		score |= 1 << 31
	} else if four {
		score |= 1 << 30
	} else if three && pairs == 1 {
		score |= 1 << 29
	} else if flush {
		score |= 1 << 28
	} else if straight {
		score |= 1 << 27
	} else if three {
		score |= 1 << 26
	} else if pairs == 2 {
		score |= 1 << 25
	} else if pairs == 1 {
		score |= 1 << 24
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
