package uno_game

import (
	"fmt"
	"math/rand"
	"time"
)

type AIUnoPlayer struct {
	*UnoPlayer
}

func NewAIUnoPlayer() *AIUnoPlayer {
	return &AIUnoPlayer{
		UnoPlayer: NewUnoPlayerData(),
	}
}

var usedAIName map[string]struct{} = make(map[string]struct{})

func (a *AIUnoPlayer) NameSelf() {
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

func (a *AIUnoPlayer) ChooseHand() *UnoCard {
	resultIdx := 0
	for idx := range a.matchedList {
		resultIdx = idx
		break
	}
	return a.ExtractCard(resultIdx)
}
