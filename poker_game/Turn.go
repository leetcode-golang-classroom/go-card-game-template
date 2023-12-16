package poker_game

import "fmt"

type Turn struct {
	winner      IPokerPlayer
	playerCards map[int]*PokerCard
	playerData  map[int]IPokerPlayer
}

func NewTurn() *Turn {
	return &Turn{
		winner:      nil,
		playerCards: make(map[int]*PokerCard),
		playerData:  make(map[int]IPokerPlayer),
	}
}
func (t *Turn) AddShow(card *PokerCard, player IPokerPlayer) {
	t.playerCards[player.GetId()] = card
	t.playerData[player.GetId()] = player
}
func (t *Turn) FindWinner() IPokerPlayer {
	var maxCard *PokerCard
	var winnerId int
	for name, card := range t.playerCards {
		if maxCard == nil {
			maxCard = card
			winnerId = name
			continue
		}
		if maxCard.Compare(card) {
			maxCard = card
			winnerId = name
		}
	}
	t.winner = t.playerData[winnerId]
	return t.winner
}

func (t *Turn) IsTurnEnd() bool {
	return len(t.playerData) == 4
}

func (t *Turn) String() string {
	result := ""
	for idx, card := range t.playerCards {
		result += fmt.Sprintf("player:%v, card: %v\n", t.playerData[idx].GetName(), card)
	}
	result += fmt.Sprintf("\nwinner: %v, point:%v\n", t.winner.GetName(), t.winner.GetPoint())
	return result
}
