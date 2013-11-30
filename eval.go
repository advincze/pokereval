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
	countRank := make([]int, 14)
	hasRank := make([]bool, 14)
	countSuit := make([]int, 4)
	minRank := hand[0].getRank()
	for _, card := range hand {
		rank := card.getRank()
		suit := card.getSuit()
		countRank[rank] += 1
		hasRank[rank] = true
		if rank < minRank {
			minRank = rank
		}
		countSuit[suit] += 1
	}
	countRank[13] = countRank[0]

	// check for flush
	var flush bool
	for _, suit := range suits {
		if countSuit[suit] == NUM_CARDS {
			flush = true
			break
		}
	}
	// println("flush:", flush)

	// check for straight
	straight := true
	if minRank == rank_Ace && hasRank[rank_10] {
		straight = hasRank[rank_Jack] && hasRank[rank_Queen] && hasRank[rank_King]
	} else {
		// log.Printf("cardsperrank%v\n", countRank)
		for rank := minRank + 1; rank < minRank+NUM_CARDS && rank < rank_AceHigh; rank++ {
			// log.Printf("checking %v %v\n", rank, countRank[rank])
			if countRank[rank] == 0 {
				straight = false
			}
		}
	}
	// println("straight:", straight)
	// primes := []uint64{41, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}

	//check four of a kind, three of a kind and pairs
	var four bool
	var three bool
	var highRank Rank
	var pairs int
	for rank := minRank; rank < rank_AceHigh; rank++ {
		switch countRank[rank] {
		case 1:
			{
				i := uint(rank)
				if i == 0 {
					i = 13
				}
				score |= 1 << (i - 1)
			}
		case 2:
			pairs += 1
			highRank = rank
		case 3:
			three = true
			highRank = rank
		case 4:
			four = true
			highRank = rank
		}
	}

	// log.Printf("score1 : %b\n", score)

	if straight && flush {
		score |= 1 << 31
	} else if four {
		score |= 1 << 30
		// score |= uint32(highRank) << 20
	} else if three && pairs == 1 {
		score |= 1 << 29
		// score |= uint32(highRank) << 20
	} else if flush {
		score |= 1 << 28
	} else if straight {
		score |= 1 << 27
	} else if three {
		score |= 1 << 26
		// score |= uint32(highRank) << 20
	} else if pairs == 2 {
		score |= 1 << 25
	} else if pairs == 1 {
		score |= 1 << 24
	}
	score |= uint32(highRank) << 20
	// log.Printf("score2 : %b\n", score)

	return &EvalResult{
		straight: straight,
		flush:    flush,
		four:     four,
		three:    three,
		pairs:    pairs,
		score:    score,
	}
}
