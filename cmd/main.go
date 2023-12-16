package main

import (
	"github.com/leetcode-golang-classroom/go-card-game-template/card_game_template"
	"github.com/leetcode-golang-classroom/go-card-game-template/poker_game"
)

func main() {
	// unoCardGame := card_game_template.NewGame[uno_game.UnoCard](
	// 	uno_game.NewUnoGame([]uno_game.IUnoPlayer{
	// 		uno_game.NewAIUnoPlayer(),
	// 		uno_game.NewAIUnoPlayer(),
	// 		uno_game.NewAIUnoPlayer(),
	// 		uno_game.NewAIUnoPlayer(),
	// 	}),
	// )
	// unoCardGame.GameFlow()
	pokerGame := card_game_template.NewGame[poker_game.PokerCard](
		poker_game.NewPokerGame(
			[]poker_game.IPokerPlayer{
				poker_game.NewHumanPokerPlayer(),
				poker_game.NewAIPokerPlayer(),
				poker_game.NewAIPokerPlayer(),
				poker_game.NewAIPokerPlayer(),
			}))
	pokerGame.GameFlow()
}
