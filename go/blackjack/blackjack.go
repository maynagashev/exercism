package blackjack

// define map with key as string and value as int
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
	return cardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Please implement the FirstTurn function
	/*
		Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:

		Stand (S)
		Hit (H)
		Split (P)
		Automatically win (W)
		Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

		If you have a pair of aces you must always split them.
		If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
		If your cards sum up to a value within the range [17, 20] you should always stand.
		If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
		If your cards sum up to 11 or lower you should always hit.
	*/

	switch {
	// If you have a pair of aces you must always split them.
	case card1 == "ace" && card2 == "ace":
		return "P"
	//If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	case isBlackJack(card1, card2) && !isPotentialBlackJack(dealerCard):
		return "W"
	// If your cards sum up to a value within the range [17, 20] you should always stand.
	case cardValues[card1]+cardValues[card2] >= 17 && cardValues[card1]+cardValues[card2] <= 20:
		return "S"
	// If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
	case cardValues[card1]+cardValues[card2] >= 12 && cardValues[card1]+cardValues[card2] <= 16 && cardValues[dealerCard] >= 7:
		return "H"
	// If your cards sum up to 11 or lower you should always hit.
	case cardValues[card1]+cardValues[card2] <= 11:
		return "H"
	default:
		return "S"
	}
}

func isPotentialBlackJack(card string) bool {
	return card == "ace" || card == "jack" || card == "queen" || card == "king" || card == "ten"
}

func isBlackJack(card1 string, card2 string) bool {
	return cardValues[card1]+cardValues[card2] == 21
}
