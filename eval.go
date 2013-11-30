package main

const NUM_CARDS = 5

type EvalResult struct {
	straight bool
	flush    bool
	four     bool
	three    bool
	pairs    int
	score    uint64
}

func eval5(hand Hand) *EvalResult {
	// println("len of hand:", len(hand))
	score := uint64(1)

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
	primes := []uint64{41, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}

	//check four of a kind, three of a kind and pairs
	var four bool
	var three bool
	var pairs int
	for rank := minRank; rank < rank_AceHigh; rank++ {
		switch countRank[rank] {
		case 1:
			{
				score *= primes[rank]
				// log.Printf("score1 : %v, %d %d\n", rank, primes[rank], score)
			}

		case 2:
			pairs += 1
		case 3:
			three = true
		case 4:
			four = true
		}
	}

	// log.Printf("score1 : %b\n", score)

	if straight {
		if flush {
			score |= 1 << 36
		} else {
			score |= 1 << 32
		}
	} else if four {
		score |= 1 << 35
	} else if three && pairs == 1 {
		score |= 1 << 34
	} else if flush {
		score |= 1 << 33
	} else if three {
		score |= 1 << 31
	} else if pairs == 2 {
		score |= 1 << 30
	} else if pairs == 1 {
		score |= 1 << 29
	}
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
