package card_game_template

import (
	"time"
)

var usedPlayerIds = map[int]struct{}{}

type Player[C any] struct {
	IPlayer[C]
	id    int
	name  string
	Hands []*C
}
type IPlayer[C any] interface {
	ChooseHand() *C
	NameSelf()
	AddHand(card *C)
	HasFinishDraw() bool
	GetName() string
	ExtractCard(idx int) *C
	GetId() int
}

func NewPlayer[C any]() *Player[C] {
	userIdExist := true
	userId := 0
	for userIdExist {
		userId = time.Now().Nanosecond()
		_, userIdExist = usedPlayerIds[userId]
		if !userIdExist {
			usedPlayerIds[userId] = struct{}{}
		}
	}
	return &Player[C]{
		Hands: []*C{},
		id:    userId,
		name:  "",
	}
}
func (p *Player[C]) AddHand(card *C) {
	p.Hands = append(p.Hands, card)
}
func (p *Player[C]) GetName() string {
	return p.name
}
func (p *Player[C]) SetName(name string) {
	p.name = name
}
func (p *Player[C]) ExtractCard(idx int) *C {
	temp := p.Hands[idx]
	if idx < len(p.Hands)-1 {
		p.Hands = append(p.Hands[:idx], p.Hands[idx+1:]...)
	} else {
		p.Hands = append([]*C{}, p.Hands[:idx]...)
	}
	return temp
}
func (p *Player[C]) GetId() int {
	return p.id
}
