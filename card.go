//go:generate stringer -type=Suit,Rank
package deck

import "fmt"

type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String() //String() -> generate by stringer in suit_string.go
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
