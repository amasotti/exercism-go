package blackjack

var cardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	if val, ok := cardValues[card]; ok {
		return val
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	val := ParseCard(card1) + ParseCard(card2)
	dealerVal := ParseCard(dealerCard)

	switch {
	case val == 22:
		return "P"
	case val == 21 && dealerVal < 10:
		return "W"
	case val == 21, val <= 20 && val >= 17:
		return "S"
	case val <= 16 && val >= 12 && dealerVal <= 6:
		return "S"
	case val <= 16 && val >= 12, val <= 11:
		return "H"
	default:
		return "S"
	}
}
