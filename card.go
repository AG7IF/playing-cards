package playing_cards

import (
	"fmt"
)

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) LaTeX() string {
	return fmt.Sprintf("%s%s", c.Rank.LaTeX(), c.Suit.LaTeX())
}
