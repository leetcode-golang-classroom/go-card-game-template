package card_game_template

type Game[C any] struct {
	IGame[C]
}

func NewGame[C any](game IGame[C]) *Game[C] {
	return &Game[C]{
		game,
	}
}

type IGame[C any] interface {
	TakeTurn()
	DisplayWinner()
	IsGameFinished() bool
	PrepareGameStep()
	GetDeck() IDeck[C]
	GetPlayers() []IPlayer[C]
	GameFlow()
}

func (g *Game[C]) GameFlow() {
	g.start()
	g.drawCards()
	g.PrepareGameStep()
	g.runGame()
	g.DisplayWinner()
}
func (g *Game[C]) drawCards() {
	players := g.GetPlayers()
	deck := g.GetDeck()
	for !players[0].HasFinishDraw() {
		deck.DrawCards(players)
	}
}
func (g *Game[C]) runGame() {
	for !g.IsGameFinished() {
		g.TakeTurn()
	}
}

func (g *Game[C]) start() {
	players := g.GetPlayers()
	deck := g.GetDeck()
	for _, player := range players {
		player.NameSelf()
	}
	deck.Shuffle()
}
