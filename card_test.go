package deck

import "testing"

type cardtesting struct {
	cd       Card
	expected string
}

func TestString(t *testing.T) {
	ct := []cardtesting{
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
				return
			}
		})
	}
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("something went wrong in number of cards")
		return
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("something went wrong in sort")
		return
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(4))
	counter := 0
	for _, c := range cards {
		if c.Suit == Joker {
			counter++
		}
	}
	if counter != 4 {
		t.Error("expected 4 Jokers; got: ", counter)
		return
	}
}

func TestFilter(t *testing.T) {
	cards := New(Filter(func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Fliter not working!")
			return
		}
	}
}
