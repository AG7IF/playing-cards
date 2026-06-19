package playing_cards

type Card struct {
	Suit Suit
	Rank Rank
	Name string
}

func (c Card) LaTeX() []byte {
	var card []byte

	card = c.Rank.LaTeX()
	card = append(card, c.Suit.LaTeX()...)

	return card
}
