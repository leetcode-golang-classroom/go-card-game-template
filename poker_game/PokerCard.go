package poker_game

import "fmt"

type PokerCard struct {
	rank Rank
	suit Suit
}

func (c *PokerCard) Compare(card *PokerCard) bool {
	return c.rank < card.rank || (c.rank == card.rank && c.suit < card.suit)
}

var suits = []string{
	"♣",
	"♦",
	"♥",
	"♠",
}

var ranks = []string{
	"2", "3", "4", "5", "6", "7", "8", "9", "10",
	"J", "Q", "K", "A",
}

func (c *PokerCard) String() string {
	return fmt.Sprintf("[%v%v]", ranks[c.rank-2], suits[c.suit])
}

type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	J
	Q
	K
	A
)

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

var Ranks = []Rank{
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	J,
	Q,
	K,
	A,
}

var Suits = []Suit{
	Club,
	Diamond,
	Heart,
	Spade,
}
