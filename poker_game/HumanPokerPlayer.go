package poker_game

import "fmt"

type HumanPokerPlayer struct {
	*PokerPlayer
}

func NewHumanPokerPlayer() *HumanPokerPlayer {
	return &HumanPokerPlayer{
		PokerPlayer: NewPlayerData(),
	}
}

func (h *HumanPokerPlayer) NameSelf() {
	fmt.Printf("Please enter player name:")
	var name string
	fmt.Scan(&name)
	h.setName(name)
}

func (h *HumanPokerPlayer) ChooseHand() *PokerCard {
	h.sortHands()
	fmt.Printf("%v, %v", len(h.hands), h.hands)
	var idx int
	answerIsInRange := false
	for !answerIsInRange {
		fmt.Printf("%v, Choose Card to Show:", h.name)
		fmt.Scan(&idx)
		answerIsInRange = idx >= 0 && idx < len(h.hands)
	}
	return h.ExtractCard(idx)
}
