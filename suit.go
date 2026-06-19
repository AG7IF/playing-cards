package playing_cards

type Suit int

const (
	Suitless Suit = iota
	Hearts
	Clubs
	Diamonds
	Spades
	Red
	Black
	Wands
	Cups
	Swords
	Pentacles
	MajorArcana
)

// LaTeX returns the LaTeX command that will print the symbol for this card's suit.
// For the standard deck of playing cards, the commands should work fine in base LaTeX. For the Tarot suits,
// however, the following packages will need to be included:
//   - esrelation
//   - fourier
//   - linearb
//   - MnSymbol
func (s Suit) LaTeX() []byte {
	var suitCmd string

	switch s {
	case Hearts:
		suitCmd = `$\heartsuit$`
	case Clubs:
		suitCmd = `$\clubsuit$`
	case Diamonds:
		suitCmd = `$\diamondsuit$`
	case Spades:
		suitCmd = `$\spadesuit$`
	case Wands:
		suitCmd = `\restrictwand`
	case Cups:
		suitCmd = `\BPcup`
	case Swords:
		suitCmd = `\textxswup`
	case Pentacles:
		suitCmd = `\pentagram`
	default:
		suitCmd = ""
	}

	return []byte(suitCmd)
}
