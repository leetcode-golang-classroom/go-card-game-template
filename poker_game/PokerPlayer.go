package poker_game

import (
	"sort"

	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type PokerPlayer struct {
	*card_game_template.Player[PokerCard]
	point int
}

type IPokerPlayer interface {
	card_game_template.IPlayer[PokerCard]
	GetPoint() int
	GainPoint()
	GetId() int
}

func NewPokerPlayer() *PokerPlayer {
	player := card_game_template.NewPlayer[PokerCard]()
	rPlayer := &PokerPlayer{
		Player: player,
		point:  0,
	}
	return rPlayer
}
func (p *PokerPlayer) HasFinishDraw() bool {
	return len(p.Hands) == 13
}

func (p *PokerPlayer) GetPoint() int {
	return p.point
}
func (p *PokerPlayer) GainPoint() {
	p.point++
}
func (p *PokerPlayer) sortHands() {
	sort.Slice(p.Hands, func(i, j int) bool {
		return !p.Hands[i].Compare(p.Hands[j])
	})
}
