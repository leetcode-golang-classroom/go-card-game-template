package poker_game

import (
	"fmt"

	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type PokerGame struct {
	deck    IPokerDeck
	players []IPokerPlayer
	turns   int
}

type IPokerGame interface {
	card_game_template.IGame[PokerGame]
}

func NewPokerGame(players []IPokerPlayer) *PokerGame {
	return &PokerGame{
		deck:    NewPokerDeck(),
		players: players,
		turns:   0,
	}
}

func (g *PokerGame) CastToIPlayers(players []IPokerPlayer) []card_game_template.IPlayer[PokerCard] {
	iplayers := []card_game_template.IPlayer[PokerCard]{}
	for _, player := range players {
		iplayers = append(iplayers, player.(card_game_template.IPlayer[PokerCard]))
	}
	return iplayers
}

func (g *PokerGame) IsGameFinished() bool {
	return g.turns == 13
}

func (g *PokerGame) TakeTurn() {
	turn := NewTurn()
	playersMap := make(map[string]IPokerPlayer)
	for _, player := range g.players {
		turn.AddShow(player.ChooseHand(), player)
		playersMap[player.GetName()] = player
	}
	winner := turn.FindWinner()
	playersMap[winner.GetName()].GainPoint()
	g.DisplayTurn(turn)
	g.NextTurn()
}
func (g *PokerGame) DisplayTurn(turn *Turn) {
	fmt.Printf("%vth\n%v\n", g.turns, turn)
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

func (g *PokerGame) GetPlayers() []card_game_template.IPlayer[PokerCard] {
	return g.CastToIPlayers(g.players)
}

func (g *PokerGame) GetDeck() card_game_template.IDeck[PokerCard] {
	return (g.deck).(card_game_template.IDeck[PokerCard])
}

func (g *PokerGame) PerpareGameStep() {}
