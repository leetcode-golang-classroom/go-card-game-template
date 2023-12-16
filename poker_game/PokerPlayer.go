package poker_game

import (
	"sort"
	"time"

	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type PokerPlayer struct {
	id    int
	hands []*PokerCard
	point int
	name  string
}

type IPokerPlayer interface {
	card_game_template.IPlayer[PokerCard]
	GetPoint() int
	GainPoint()
	GetId() int
}

func (p *PokerPlayer) GetPoint() int {
	return p.point
}
func (p *PokerPlayer) GainPoint() {
	p.point++
}
func (p *PokerPlayer) setName(name string) {
	p.name = name
}

func (p *PokerPlayer) HasFinishDraw() bool {
	return len(p.hands) == 13
}

func (p *PokerPlayer) AddHand(card *PokerCard) {
	p.hands = append(p.hands, card)
}

func (p *PokerPlayer) GetName() string {
	return p.name
}

func (p *PokerPlayer) ExtractCard(idx int) *PokerCard {
	temp := p.hands[idx]
	if idx < len(p.hands)-1 {
		p.hands = append(p.hands[:idx], p.hands[idx+1:]...)
	} else {
		p.hands = append([]*PokerCard{}, p.hands[:idx]...)
	}
	return temp
}

func (p *PokerPlayer) sortHands() {
	sort.Slice(p.hands, func(i, j int) bool {
		return !p.hands[i].Compare(p.hands[j])
	})
}
func NewPlayerData() *PokerPlayer {
	return &PokerPlayer{
		point: 0,
		hands: []*PokerCard{},
		name:  "",
		id:    time.Now().Nanosecond(),
	}
}

func (p *PokerPlayer) GetId() int {
	return p.id
}
