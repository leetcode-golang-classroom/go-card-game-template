package poker_game

import "fmt"

type HumanPokerPlayer struct {
	*PokerPlayer
}

func NewHumanPokerPlayer() *HumanPokerPlayer {
	return &HumanPokerPlayer{
		NewPokerPlayer(),
	}
}

func (h *HumanPokerPlayer) NameSelf() {
	fmt.Printf("Please enter player name:")
	var name string
	fmt.Scan(&name)
	h.SetName(name)
}

func (h *HumanPokerPlayer) ChooseHand() *PokerCard {
	h.sortHands()
	fmt.Printf("%v, points: %v\ntotal:%v cards, hands:%v\n", h.GetName(), h.point, len(h.Hands), h.Hands)
	var idx int
	answerIsInRange := false
	for !answerIsInRange {
		fmt.Printf("%v, Choose Card to Show:", h.GetName())
		fmt.Scanf("%d", &idx)
		answerIsInRange = idx >= 0 && idx < len(h.Hands)
	}
	return h.ExtractCard(idx)
}
