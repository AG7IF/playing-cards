package playing_cards

type Deck []Card

var newDeck = []Card{
	{Suit: Hearts, Rank: Ace},
	{Suit: Hearts, Rank: Duce},
	{Suit: Hearts, Rank: Three},
	{Suit: Hearts, Rank: Four},
	{Suit: Hearts, Rank: Five},
	{Suit: Hearts, Rank: Six},
	{Suit: Hearts, Rank: Seven},
	{Suit: Hearts, Rank: Eight},
	{Suit: Hearts, Rank: Nine},
	{Suit: Hearts, Rank: Ten},
	{Suit: Hearts, Rank: Jack},
	{Suit: Hearts, Rank: Queen},
	{Suit: Hearts, Rank: King},

	{Suit: Clubs, Rank: Ace},
	{Suit: Clubs, Rank: Duce},
	{Suit: Clubs, Rank: Three},
	{Suit: Clubs, Rank: Four},
	{Suit: Clubs, Rank: Five},
	{Suit: Clubs, Rank: Six},
	{Suit: Clubs, Rank: Seven},
	{Suit: Clubs, Rank: Eight},
	{Suit: Clubs, Rank: Nine},
	{Suit: Clubs, Rank: Ten},
	{Suit: Clubs, Rank: Jack},
	{Suit: Clubs, Rank: Queen},
	{Suit: Clubs, Rank: King},

	{Suit: Diamonds, Rank: Ace},
	{Suit: Diamonds, Rank: Duce},
	{Suit: Diamonds, Rank: Three},
	{Suit: Diamonds, Rank: Four},
	{Suit: Diamonds, Rank: Five},
	{Suit: Diamonds, Rank: Six},
	{Suit: Diamonds, Rank: Seven},
	{Suit: Diamonds, Rank: Eight},
	{Suit: Diamonds, Rank: Nine},
	{Suit: Diamonds, Rank: Ten},
	{Suit: Diamonds, Rank: Jack},
	{Suit: Diamonds, Rank: Queen},
	{Suit: Diamonds, Rank: King},

	{Suit: Spades, Rank: Ace},
	{Suit: Spades, Rank: Duce},
	{Suit: Spades, Rank: Three},
	{Suit: Spades, Rank: Four},
	{Suit: Spades, Rank: Five},
	{Suit: Spades, Rank: Six},
	{Suit: Spades, Rank: Seven},
	{Suit: Spades, Rank: Eight},
	{Suit: Spades, Rank: Nine},
	{Suit: Spades, Rank: Ten},
	{Suit: Spades, Rank: Jack},
	{Suit: Spades, Rank: Queen},
	{Suit: Spades, Rank: King},

	{Suit: Black, Rank: Joker},
	{Suit: Red, Rank: Joker},
}

func NewDeck() Deck {
	var d Deck

	for _, c := range newDeck {
		d = append(d, c)
	}

	return d
}
