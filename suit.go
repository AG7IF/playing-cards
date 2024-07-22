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
)

func (s Suit) LaTeX() string {
	switch s {
	case Hearts:
		return `$\heartsuit$`
	case Clubs:
		return `$\clubsuit$`
	case Diamonds:
		return `$\diamondsuit$`
	case Spades:
		return `$\spadesuit$`
	default:
		return ""
	}
}
