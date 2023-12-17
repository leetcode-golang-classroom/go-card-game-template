package uno_game

import "fmt"

type HumanUnoPlayer struct {
	*UnoPlayer
}

func NewHumanUnoPlayer() *HumanUnoPlayer {
	return &HumanUnoPlayer{
		UnoPlayer: NewUnoPlayerData(),
	}
}

func (h *HumanUnoPlayer) NameSelf() {
	fmt.Printf("Please enter player name:")
	var name string
	fmt.Scan(&name)
	h.setName(name)
}

func (h *HumanUnoPlayer) ChooseHand() *UnoCard {
	fmt.Printf("human player: %v\ntotal %v cards in hand, total hands: %v\ntotal %v card could use, usable cards: %v\n", h.name, len(h.hands), h.hands, len(h.matchedList), h.matchedList)
	var idx int
	validChoice := false
	for !validChoice {
		fmt.Printf("%v, please choose card to show:", h.GetName())
		fmt.Scan(&idx)
		_, validChoice = h.matchedList[idx]
	}
	return h.ExtractCard(idx)
}
