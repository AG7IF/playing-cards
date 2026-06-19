package playing_cards

import (
	"strconv"
)

type Rank int

const (
	Joker Rank = iota
	Ace
	Duce
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Knight
	Queen
	King
)

func (r Rank) String() string {
	switch r {
	case Joker:
		return "JK"
	case Ace:
		return "A"
	case Jack:
		return "J"
	case Knight:
		return "C"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return strconv.Itoa(int(r))
	}
}

func (r Rank) LaTeX() []byte {
	return []byte(r.String())
}
