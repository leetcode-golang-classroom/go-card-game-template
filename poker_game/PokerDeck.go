package poker_game

import (
	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type PokerDeck struct {
	*card_game_template.Deck[PokerCard]
}
type IPokerDeck interface {
	card_game_template.IDeck[PokerCard]
}

func NewPokerDeck() IPokerDeck {
	deck := &PokerDeck{
		card_game_template.NewDeck[PokerCard](),
	}
	deck.InitData()
	return deck
}
func (d *PokerDeck) InitData() {
	for _, suit := range Suits {
		for _, rank := range Ranks {
			d.Cards = append(d.Cards, &PokerCard{rank: rank, suit: suit})
		}
	}
}
