package card_game_template

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck[C any] struct {
	IDeck[C]
	Cards []*C
}
type IDeck[C any] interface {
	Shuffle()
	DrawCards(players []IPlayer[C])
	DrawCard() *C
	InitData()
	ShowDeck()
}

func NewDeck[C any]() *Deck[C] {
	return &Deck[C]{
		Cards: []*C{},
	}
}
func (d *Deck[C]) Shuffle() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
func (d *Deck[C]) DrawCards(players []IPlayer[C]) {
	for _, player := range players {
		card := d.DrawCard()
		player.AddHand(card)
	}
}
func (d *Deck[C]) DrawCard() *C {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := rand.Intn(len(d.Cards))
	temp := d.Cards[idx]
	if idx < len(d.Cards)-1 {
		d.Cards = append(d.Cards[:idx], d.Cards[idx+1:]...)
	} else {
		d.Cards = append([]*C{}, d.Cards[:idx]...)
	}
	return temp
}
func (d *Deck[C]) ShowDeck() {
	fmt.Printf("%v", d.Cards)
}
