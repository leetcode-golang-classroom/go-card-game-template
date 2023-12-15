package poker_game

import (
	"math/rand"
	"time"

	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type PokerDeck struct {
	cards []*PokerCard
}
type IPokerDeck interface {
	card_game_template.IDeck[PokerCard]
}

func NewPokerDeck() *PokerDeck {
	deck := &PokerDeck{
		cards: []*PokerCard{},
	}
	deck.InitData()
	return deck
}
func (d *PokerDeck) InitData() {
	d.cards = []*PokerCard{}
	for _, suit := range Suits {
		for _, rank := range Ranks {
			d.cards = append(d.cards, &PokerCard{rank: rank, suit: suit})
		}
	}
}

func (d *PokerDeck) Shuffle() {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func (d *PokerDeck) DrawCards(players []card_game_template.IPlayer[PokerCard]) {
	for _, player := range players {
		card := d.DrawCard()
		player.AddHand(card)
	}
}

func (d *PokerDeck) DrawCard() *PokerCard {
	rand.NewSource(time.Now().UnixNano())
	idx := rand.Intn(len(d.cards))
	temp := d.cards[idx]
	if idx < len(d.cards)-1 {
		d.cards = append(d.cards[:idx], d.cards[idx+1:]...)
	} else {
		d.cards = append([]*PokerCard{}, d.cards[:idx]...)
	}
	return temp
}
