package main

import (
	"testing"
)

func TestShouldRecogniseFlush(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_2),
		NewCard(clubs, rank_4),
		NewCard(clubs, rank_6),
		NewCard(clubs, rank_8),
		NewCard(clubs, rank_10),
	})
	if !eval5(hand).flush {
		t.Errorf("hand: %v should be recognised as a flush", Cards(hand))
	}
}

func TestShouldRecogniseNonFlush(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_2),
		NewCard(clubs, rank_4),
		NewCard(clubs, rank_6),
		NewCard(clubs, rank_8),
		NewCard(diamonds, rank_10),
	})
	if eval5(hand).flush {
		t.Errorf("hand: %v should not be recognised as a flush", Cards(hand))
	}
}

func TestShouldRecogniseStraightWithNoAce(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_2),
		NewCard(clubs, rank_4),
		NewCard(clubs, rank_5),
		NewCard(diamonds, rank_3),
		NewCard(clubs, rank_6),
	})
	if !eval5(hand).straight {
		t.Errorf("hand: %v should be recognised as a straight", Cards(hand))
	}
}

func TestShouldRecogniseStraightWithAceAsFirstCard(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_2),
		NewCard(spades, rank_4),
		NewCard(clubs, rank_5),
		NewCard(clubs, rank_Ace),
		NewCard(diamonds, rank_3),
	})
	if !eval5(hand).straight {
		t.Errorf("hand: %v should be recognised as a straight", Cards(hand))
	}
}

func TestShouldRecogniseStraightWithAceAsLastCard(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_Queen),
		NewCard(spades, rank_10),
		NewCard(clubs, rank_Jack),
		NewCard(clubs, rank_Ace),
		NewCard(diamonds, rank_King),
	})
	if !eval5(hand).straight {
		t.Errorf("hand: %v should be recognised as a straight", Cards(hand))
	}
}

func TestShouldRecogniseNonStraightWithAce(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_Queen),
		NewCard(spades, rank_9),
		NewCard(clubs, rank_Jack),
		NewCard(clubs, rank_Ace),
		NewCard(diamonds, rank_King),
	})
	if eval5(hand).straight {
		t.Errorf("hand: %v should not be recognised as a straight", Cards(hand))
	}
}

func TestShouldRecogniseFourOfAKind(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_2),
		NewCard(spades, rank_2),
		NewCard(diamonds, rank_2),
		NewCard(hearts, rank_2),
		NewCard(diamonds, rank_3),
	})
	if !eval5(hand).four {
		t.Errorf("hand: %v should be recognised as a four of a kind", Cards(hand))
	}
}

func TestShouldRecogniseFourOfAKindAce(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_Ace),
		NewCard(spades, rank_Ace),
		NewCard(diamonds, rank_Ace),
		NewCard(hearts, rank_Ace),
		NewCard(diamonds, rank_2),
	})
	if !eval5(hand).four {
		t.Errorf("hand: %v should be recognised as a four of a kind", Cards(hand))
	}
}

func TestShouldRecogniseThreeOfAKind(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_King),
		NewCard(spades, rank_King),
		NewCard(diamonds, rank_King),
		NewCard(hearts, rank_Ace),
		NewCard(diamonds, rank_2),
	})
	if !eval5(hand).three {
		t.Errorf("hand: %v should be recognised as a three of a kind", Cards(hand))
	}
}

func TestShouldRecogniseFullHouse(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_King),
		NewCard(spades, rank_King),
		NewCard(diamonds, rank_Ace),
		NewCard(hearts, rank_Ace),
		NewCard(spades, rank_Ace),
	})
	if res := eval5(hand); !(res.three && res.pairs == 1) {
		t.Errorf("hand: %v should be recognised as full house : three: %v and one pair: %d", Cards(hand), res.three, res.pairs)
	}
}

func TestShouldRecogniseTwoPairs(t *testing.T) {
	hand := Hand([]Card{
		NewCard(clubs, rank_King),
		NewCard(spades, rank_King),
		NewCard(diamonds, rank_Ace),
		NewCard(hearts, rank_Ace),
		NewCard(diamonds, rank_2),
	})
	if res := eval5(hand); res.pairs != 2 {
		t.Errorf("hand: %v should be recognised as two pairs, were: %d", Cards(hand), res.pairs)
	}
}

func TestShouldScoreHigherWithHigherHighcard(t *testing.T) {
	lowerHand := Hand(ParseCards("C2", "S7", "D3", "H4", "S5"))
	higherHand := Hand(ParseCards("C9", "S4", "D5", "H2", "D6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithHigherPairButLowerRanks(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S7", "D3", "H4", "DK"))
	higherHand := Hand(ParseCards("HQ", "DQ", "S7", "H3", "D4"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithHigherHighcardAce(t *testing.T) {
	lowerHand := Hand(ParseCards("C5", "S7", "D3", "H4", "D2"))
	higherHand := Hand(ParseCards("CA", "S4", "D5", "H2", "D6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithOnePairHigherKicker(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S7", "D3", "H4", "D5"))
	higherHand := Hand(ParseCards("H7", "D7", "S3", "H4", "D6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithOnePairHigherKickerAce(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S7", "D3", "H4", "DK"))
	higherHand := Hand(ParseCards("H7", "D7", "S3", "H4", "DA"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithOnePairVsHighcard(t *testing.T) {
	lowerHand := Hand(ParseCards("C6", "S7", "D3", "H4", "D2"))
	higherHand := Hand(ParseCards("C2", "S4", "D5", "H2", "D6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithOnePairVsHighcardAce(t *testing.T) {
	lowerHand := Hand(ParseCards("C6", "SA", "D3", "H4", "D2"))
	higherHand := Hand(ParseCards("C2", "S4", "D5", "H2", "D6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithTwoPairsVsOnePair(t *testing.T) {
	lowerHand := Hand(ParseCards("CK", "SK", "D9", "H8", "D7"))
	higherHand := Hand(ParseCards("C2", "S3", "D3", "H2", "D6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithTwoPairs(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S7", "C3", "S3", "DQ"))
	higherHand := Hand(ParseCards("H7", "D7", "HQ", "DQ", "D3"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithTwoPairsLowerPairDecides(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S7", "C3", "S3", "DQ"))
	higherHand := Hand(ParseCards("H7", "D7", "H4", "D4", "D3"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithTwoPairsHigherKicker(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S7", "C3", "S3", "DQ"))
	higherHand := Hand(ParseCards("H7", "D7", "H3", "D3", "DK"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithTwoPairsHigherKickerAce(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S7", "C3", "S3", "DK"))
	higherHand := Hand(ParseCards("H7", "D7", "H3", "D3", "DA"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWith3oakVsTwoPairs(t *testing.T) {
	lowerHand := Hand(ParseCards("CK", "SK", "DQ", "HQ", "D7"))
	higherHand := Hand(ParseCards("C2", "S2", "D3", "H2", "D6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWith3oakHigherKicker(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S7", "D7", "S3", "DQ"))
	higherHand := Hand(ParseCards("C7", "S7", "D7", "S3", "DK"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithStraightVs3oak(t *testing.T) {
	lowerHand := Hand(ParseCards("CK", "SK", "DK", "HQ", "D7"))
	higherHand := Hand(ParseCards("C2", "S3", "D4", "H5", "D6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithHigherStraight(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S8", "D9", "S10", "DJ"))
	higherHand := Hand(ParseCards("C8", "S9", "D10", "SJ", "DQ"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithHigherStraightWithAce(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "S8", "D9", "S10", "DJ"))
	higherHand := Hand(ParseCards("CA", "SK", "D10", "SJ", "DQ"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithFlushVsStraight(t *testing.T) {
	lowerHand := Hand(ParseCards("SK", "SQ", "DJ", "H9", "D10"))
	higherHand := Hand(ParseCards("C2", "C3", "C4", "C5", "C7"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithFlushVsFlushHigherCads(t *testing.T) {
	lowerHand := Hand(ParseCards("C7", "C8", "C9", "C10", "CQ"))
	higherHand := Hand(ParseCards("S2", "S3", "S4", "S5", "SK"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithFullhouseVsFlush(t *testing.T) {
	lowerHand := Hand(ParseCards("SK", "SQ", "SJ", "SA", "S9"))
	higherHand := Hand(ParseCards("C2", "D2", "H3", "D3", "S3"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}
func TestShouldScoreHigherWithFUllHouseVsFullHouse(t *testing.T) {
	lowerHand := Hand(ParseCards("C8", "D8", "C9", "D9", "S9"))
	higherHand := Hand(ParseCards("S8", "D8", "S10", "D10", "H10"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWith4oakVsFullHouse(t *testing.T) {
	lowerHand := Hand(ParseCards("CK", "DK", "HQ", "DQ", "SQ"))
	higherHand := Hand(ParseCards("C2", "D2", "S2", "H2", "S3"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWith4oakVs4oak(t *testing.T) {
	lowerHand := Hand(ParseCards("C8", "D8", "S8", "H8", "D9"))
	higherHand := Hand(ParseCards("C9", "D9", "S9", "H9", "D8"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWith4oakVs4oakDecideOnKickerCard(t *testing.T) {
	lowerHand := Hand(ParseCards("C9", "D9", "S9", "H9", "DJ"))
	higherHand := Hand(ParseCards("C9", "D9", "S9", "H9", "DQ"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithStraightFlushVs4oak(t *testing.T) {
	lowerHand := Hand(ParseCards("CK", "DK", "HK", "SK", "SQ"))
	higherHand := Hand(ParseCards("C2", "C3", "C4", "C5", "C6"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func TestShouldScoreHigherWithStraightFlushVsStraightFlush(t *testing.T) {
	lowerHand := Hand(ParseCards("C8", "C9", "C10", "CJ", "CQ"))
	higherHand := Hand(ParseCards("C9", "C10", "CJ", "CQ", "CK"))

	ShouldScoreHigherWithBetterCombination(t, lowerHand, higherHand)
}

func ShouldScoreHigherWithBetterCombination(t *testing.T, lowerHand, higherHand Hand) {
	if lowerRes, higherRes := eval5(lowerHand), eval5(higherHand); lowerRes.score >= higherRes.score {
		t.Errorf("hand: %v with res: %#v should score higher than %v res: %#v", Cards(higherHand), higherRes, Cards(lowerHand), lowerRes)
	}
}
