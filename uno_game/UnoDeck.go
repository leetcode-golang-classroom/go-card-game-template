package uno_game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type UnoDeck struct {
	cards []*UnoCard
}
type IUnoDeck interface {
	card_game_template.IDeck[UnoCard]
	ShowDeck()
	AddCards(card []*UnoCard)
	IsDeckEmpty() bool
}

func NewUnoDeck() *UnoDeck {
	deck := &UnoDeck{
		cards: []*UnoCard{},
	}
	deck.InitData()
	return deck
}

func (d *UnoDeck) InitData() {
	d.cards = []*UnoCard{}
	for _, number := range Numbers {
		for _, color := range Colors {
			d.cards = append(d.cards, &UnoCard{number: number, color: color})
		}
	}
}

func (d *UnoDeck) Shuffle() {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *UnoDeck) DrawCards(players []card_game_template.IPlayer[UnoCard]) {
	for _, player := range players {
		card := d.DrawCard()
		player.AddHand(card)
	}
}

func (d *UnoDeck) DrawCard() *UnoCard {
	rand.NewSource(time.Now().UnixNano())
	idx := rand.Intn(len(d.cards))
	temp := d.cards[idx]
	if idx < len(d.cards)-1 {
		d.cards = append(d.cards[:idx], d.cards[idx+1:]...)
	} else {
		d.cards = append([]*UnoCard{}, d.cards[:idx]...)
	}
	return temp
}

func (d *UnoDeck) AddCards(cards []*UnoCard) {
	d.cards = append(d.cards, cards...)
}

func (d *UnoDeck) ShowDeck() {
	fmt.Print("deck:")
	for _, card := range d.cards {
		fmt.Printf("%v,", card)
	}
	fmt.Println()
}

func (d *UnoDeck) IsDeckEmpty() bool {
	return len(d.cards) == 0
}
