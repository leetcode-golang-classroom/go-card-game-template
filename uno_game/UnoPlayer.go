package uno_game

import (
	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type UnoPlayer struct {
	*card_game_template.Player[UnoCard]
	matchedList map[int]*UnoCard
}

type IUnoPlayer interface {
	card_game_template.IPlayer[UnoCard]
	FindMatchedHand(card *UnoCard) map[int]*UnoCard
	IsHandEmpty() bool
	ClearMatched()
}

func (p *UnoPlayer) IsHandEmpty() bool {
	return len(p.Hands) == 0
}

func (p *UnoPlayer) HasFinishDraw() bool {
	return len(p.Hands) == 5
}
func (p *UnoPlayer) FindMatchedHand(card *UnoCard) map[int]*UnoCard {
	matchedCards := make(map[int]*UnoCard)
	for idx, hand := range p.Hands {
		if hand.Equal(card) {
			matchedCards[idx] = hand
		}
	}
	p.matchedList = matchedCards
	return matchedCards
}
func (p *UnoPlayer) ClearMatched() {
	p.matchedList = make(map[int]*UnoCard)
}
func NewUnoPlayerData() *UnoPlayer {
	player := card_game_template.NewPlayer[UnoCard]()
	return &UnoPlayer{
		Player:      player,
		matchedList: make(map[int]*UnoCard),
	}
}
