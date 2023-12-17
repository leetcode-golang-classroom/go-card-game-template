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
		if !nameExist {
			usedAIName[name] = struct{}{}
		}
	}
	a.setName(name)
}

func (a *AIPokerPlayer) ChooseHand() *PokerCard {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	a.sortHands()
	// rand from 0,1,2,3, len(a.hands)-1
	lenHands := len(a.hands) - 1
	// randIdx := 0
	idx := 0
	if lenHands > 4 {
		randList := []int{0, 1, 2, 3, lenHands}
		randIdx := rand.Intn(len(randList))
		// fmt.Printf("randIdx: %v", randIdx)
		idx = randList[randIdx]
	} else {
		idx = rand.Intn(len(a.hands))
	}

	return a.ExtractCard(idx)
}
