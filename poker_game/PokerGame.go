package poker_game

import (
	"fmt"

	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type PokerGame struct {
	*card_game_template.Game[PokerCard]
	deck    IPokerDeck
	players []IPokerPlayer
	turns   int
}

func NewPokerGame(players []IPokerPlayer) card_game_template.IGame[PokerCard] {
	game := &PokerGame{
		deck:    NewPokerDeck(),
		players: players,
		turns:   0,
	}
	return card_game_template.NewGame[PokerCard](game)
}

func (g *PokerGame) IsGameFinished() bool {
	return g.turns == 13
}

func (g *PokerGame) TakeTurn() {
	turn := NewTurn()
	playersMap := make(map[int]IPokerPlayer)
	for _, player := range g.players {
		turn.AddShow(player.ChooseHand(), player)
		playersMap[player.GetId()] = player
	}
	winner := turn.FindWinner()
	playersMap[winner.GetId()].GainPoint()
	g.DisplayTurn(turn)
	g.NextTurn()
}
func (g *PokerGame) DisplayTurn(turn *Turn) {
	fmt.Printf("======\n%vth\n%v\n=========\n", g.turns, turn)
}
func (g *PokerGame) NextTurn() {
	g.turns++
}
func (g *PokerGame) DisplayWinner() {
	winners := []IPokerPlayer{}
	for _, player := range g.players {
		if len(winners) == 0 {
			winners = append(winners, player)
			continue
		}
		if winners[0].GetPoint() < player.GetPoint() {
			winners = append([]IPokerPlayer{}, player)
		} else if winners[0].GetPoint() == player.GetPoint() {
			winners = append(winners, player)
		}
	}
	fmt.Print("winners are ")
	for idx, winner := range winners {
		fmt.Printf("%v", winner.GetName())
		if idx != len(winners)-1 {
			fmt.Printf(",")
		} else {
			fmt.Printf(", points %v\n", winner.GetPoint())
		}
	}
}

func (g *PokerGame) PrepareGameStep() {}
func (g *PokerGame) GetDeck() card_game_template.IDeck[PokerCard] {
	return g.deck.(card_game_template.IDeck[PokerCard])
}
func (g *PokerGame) GetPlayers() []card_game_template.IPlayer[PokerCard] {
	players := []card_game_template.IPlayer[PokerCard]{}
	for _, player := range g.players {
		players = append(players, player.(card_game_template.IPlayer[PokerCard]))
	}
	return players
}
