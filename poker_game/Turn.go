package poker_game

import "fmt"

type Turn struct {
	winner      IPokerPlayer
	playerCards map[string]*PokerCard
	playerData  map[string]IPokerPlayer
}

func NewTurn() *Turn {
	return &Turn{
		winner:      nil,
		playerCards: make(map[string]*PokerCard),
		playerData:  make(map[string]IPokerPlayer),
	}
}
func (t *Turn) AddShow(card *PokerCard, player IPokerPlayer) {
	t.playerCards[player.GetName()] = card
	t.playerData[player.GetName()] = player
}
func (t *Turn) FindWinner() IPokerPlayer {
	var maxCard *PokerCard
	var winnerName string
	for name, card := range t.playerCards {
		if maxCard == nil {
			maxCard = card
			winnerName = name
			continue
		}
		if maxCard.Compare(card) {
			maxCard = card
			winnerName = name
		}
	}
	t.winner = t.playerData[winnerName]
	return t.winner
}

func (t *Turn) IsTurnEnd() bool {
	return len(t.playerData) == 4
}

func (t *Turn) String() string {
	result := ""
	for name, card := range t.playerCards {
		result += fmt.Sprintf("player:%v, card: %v\n", name, card)
	}
	result += fmt.Sprintf("\nwinner: %v, point:%v\n", t.winner.GetName(), t.winner.GetPoint())
	return result
}
