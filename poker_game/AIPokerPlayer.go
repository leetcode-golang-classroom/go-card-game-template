package poker_game

import (
	"fmt"
	"math/rand"
	"time"
)

type AIPokerPlayer struct {
	*PokerPlayer
}

var usedAIName map[string]struct{} = make(map[string]struct{})

func NewAIPokerPlayer() *AIPokerPlayer {
	return &AIPokerPlayer{
		PokerPlayer: NewPlayerData(),
	}
}

func (a *AIPokerPlayer) NameSelf() {
	nameExist := false
	name := ""
	for !nameExist {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		number := rand.Intn(1000)
		name = fmt.Sprintf("AI-Number%v", number)
		_, nameExist = usedAIName[name]
		usedAIName[name] = struct{}{}
	}
	a.setName(name)
}

func (a *AIPokerPlayer) ChooseHand() *PokerCard {
	a.sortHands()
	idx := 0
	return a.ExtractCard(idx)
}
