package deck

import "testing"

type cardtesting struct {
	cd Card
	expected string
}

func TestString(t *testing.T) {	
	ct := []cardtesting {
		{Card{Rank: Ace, Suit: Heart}, "Ace of Hearts"},
		{Card{Rank: Two, Suit: Club}, "Two of Clubs"},
		{Card{Rank: Jack, Suit: Diamond}, "Jack of Diamonds"},
		{Card{Rank: Ten, Suit: Spade}, "Ten of Spades"},
		{Card{Rank: Queen, Suit: Club}, "Queen of Clubs"},
		{Card{Rank: King, Suit: Heart}, "King of Hearts"},
		{Card{Suit: Joker}, "Joker"},
	}

	for _, tt := range ct {
		t.Run(tt.cd.Rank.String(), func(t *testing.T) {
			if got := tt.cd.String(); got != tt.expected {
				t.Errorf("expected :%s ---- got:%s ", tt.expected, got)
			}
		})
	}
}