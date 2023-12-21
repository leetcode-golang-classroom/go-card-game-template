package uno_game

import (
	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type UnoDeck struct {
	*card_game_template.Deck[UnoCard]
}
type IUnoDeck interface {
	card_game_template.IDeck[UnoCard]
	ShowDeck()
	AddCards(card []*UnoCard)
	IsDeckEmpty() bool
}

func NewUnoDeck() *UnoDeck {
	deck := &UnoDeck{
		card_game_template.NewDeck[UnoCard](),
	}
	deck.InitData()
	return deck
}

func (d *UnoDeck) InitData() {
	d.Cards = []*UnoCard{}
	for _, number := range Numbers {
		for _, color := range Colors {
			d.Cards = append(d.Cards, &UnoCard{number: number, color: color})
		}
	}
}

func (d *UnoDeck) AddCards(cards []*UnoCard) {
	d.Cards = append(d.Cards, cards...)
}

func (d *UnoDeck) IsDeckEmpty() bool {
	return len(d.Cards) == 0
}
