package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Parse := ParseCard(card1)
	card2Parse := ParseCard(card2)
	dealerParse := ParseCard(dealerCard)
	cardSum := card1Parse + card2Parse

	switch {
	case cardSum == 22:
		return "P"
	case cardSum == 21:
		if dealerParse == 11 || dealerParse == 10 {
			return "S"
		}
		return "W"
	case cardSum >= 17 && cardSum <= 20:
		return "S"
	case cardSum >= 12 && cardSum <= 16:
		if dealerParse < 7 {
			return "S"
		}
		return "H"
	case cardSum <= 11:
		return "H"
	}
	panic("Please implement the FirstTurn function")
}
