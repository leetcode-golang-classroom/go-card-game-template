package uno_game

import (
	"time"

	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type UnoPlayer struct {
	id          int
	hands       []*UnoCard
	name        string
	matchedList map[int]*UnoCard
}

type IUnoPlayer interface {
	card_game_template.IPlayer[UnoCard]
	FindMatchedHand(card *UnoCard) map[int]*UnoCard
	IsHandEmpty() bool
	ClearMatched()
}

func (p *UnoPlayer) IsHandEmpty() bool {
	return len(p.hands) == 0
}
func (p *UnoPlayer) GetName() string {
	return p.name
}
func (p *UnoPlayer) setName(name string) {
	p.name = name
}

func (p *UnoPlayer) AddHand(card *UnoCard) {
	p.hands = append(p.hands, card)
}

func (p *UnoPlayer) HasFinishDraw() bool {
	return len(p.hands) == 5
}
func (p *UnoPlayer) FindMatchedHand(card *UnoCard) map[int]*UnoCard {
	matchedCards := make(map[int]*UnoCard)
	for idx, hand := range p.hands {
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
	return &UnoPlayer{
		hands:       []*UnoCard{},
		name:        "",
		matchedList: make(map[int]*UnoCard),
		id:          time.Now().Nanosecond(),
	}
}
func (p *UnoPlayer) ExtractCard(idx int) *UnoCard {
	temp := p.hands[idx]
	if idx < len(p.hands)-1 {
		p.hands = append(p.hands[:idx], p.hands[idx+1:]...)
	} else {
		p.hands = append([]*UnoCard{}, p.hands[:idx]...)
	}
	return temp
}
