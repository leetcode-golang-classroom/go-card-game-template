package main

import (
	"github.com/leetcode-golang-classroom/go-card-game-template/uno_game"
)

func main() {
	unoCardGame := uno_game.NewUnoGame([]uno_game.IUnoPlayer{
		uno_game.NewHumanUnoPlayer(),
		uno_game.NewAIUnoPlayer(),
		uno_game.NewAIUnoPlayer(),
		uno_game.NewAIUnoPlayer(),
	})
	unoCardGame.GameFlow()
	// pokerGame := poker_game.NewPokerGame([]poker_game.IPokerPlayer{
	// 	poker_game.NewHumanPokerPlayer(),
	// 	poker_game.NewAIPokerPlayer(),
	// 	poker_game.NewAIPokerPlayer(),
	// 	poker_game.NewAIPokerPlayer(),
	// })
	// pokerGame.GameFlow()

}
