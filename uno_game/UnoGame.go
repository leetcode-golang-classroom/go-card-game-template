package uno_game

import (
	"fmt"

	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
)

type UnoGame struct {
	*card_game_template.Game[UnoCard]
	deck    IUnoDeck
	players []IUnoPlayer
	turns   int
	table   *Table
}

type IUnoGame interface {
	card_game_template.IGame[UnoCard]
}

func NewUnoGame(players []IUnoPlayer) *UnoGame {
	return &UnoGame{
		deck:    NewUnoDeck(),
		players: players,
		turns:   0,
		table:   NewTable(),
	}
}

func (g *UnoGame) CastToIPlayers(players []IUnoPlayer) []card_game_template.IPlayer[UnoCard] {
	iplayers := []card_game_template.IPlayer[UnoCard]{}
	for _, player := range g.players {
		iplayers = append(iplayers, player.(card_game_template.IPlayer[UnoCard]))
	}
	return iplayers
}

func (g *UnoGame) IsGameFinished() bool {
	return g.players[g.GetTurn()].IsHandEmpty()
}

func (g *UnoGame) GetTurn() int {
	return g.turns % len(g.players)
}

func (g *UnoGame) NextTurn() {
	g.turns++
}
func (g *UnoGame) handleDeckAddHand(player IUnoPlayer) {
	card := g.deck.DrawCard()
	player.AddHand(card)
}
func (g *UnoGame) handleNoMatched(player IUnoPlayer) {
	fmt.Printf("player %v, handle not matched\n", player.GetName())
	if g.deck.IsDeckEmpty() {
		cards := g.table.PushBackToDeck()
		g.deck.AddCards(cards)
		return
	}
	g.handleDeckAddHand(player)
}
func (g *UnoGame) PrepareGameStep() {
	card := g.deck.DrawCard()
	g.table.AddCard(card)
}
func (g *UnoGame) TakeTurn() {
	currentPlayer := g.players[g.GetTurn()]
	topTableCard := g.table.FindTopest()
	matchedList := currentPlayer.FindMatchedHand(topTableCard)
	if len(matchedList) == 0 {
		g.handleNoMatched(currentPlayer)
		g.NextTurn()
		return
	}
	fmt.Printf("========\n%v turn start\n", g.turns+1)
	g.deck.ShowDeck()
	fmt.Printf("table top: %v\n", g.table.FindTopest())
	card := currentPlayer.ChooseHand()
	fmt.Printf("player %v, show %v\n", currentPlayer.GetName(), card)
	g.table.AddCard(card)
	currentPlayer.ClearMatched()
	fmt.Printf("%v turn end\n", g.turns+1)
	if !g.IsGameFinished() {
		g.NextTurn()
	}
}

func (g *UnoGame) DisplayWinner() {
	fmt.Printf("winner is %v\n", g.players[g.GetTurn()].GetName())
}

func (g *UnoGame) GetPlayers() []card_game_template.IPlayer[UnoCard] {
	return g.CastToIPlayers(g.players)
}

func (g *UnoGame) GetDeck() card_game_template.IDeck[UnoCard] {
	return (g.deck).(card_game_template.IDeck[UnoCard])
}
