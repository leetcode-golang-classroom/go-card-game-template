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
	GetPlayers() []IPlayer[C]
	GetDeck() IDeck[C]
	PerpareGameStep()
}

func (g *Game[C]) GameFlow() {
	g.start()
	g.drawCards()
	g.PerpareGameStep()
	g.runGame()
	g.DisplayWinner()
}
func (g *Game[C]) drawCards() {
	iplayers := g.GetPlayers()
	iDeck := g.GetDeck()
	for !iplayers[0].HasFinishDraw() {
		iDeck.DrawCards(iplayers)
	}
}
func (g *Game[C]) runGame() {
	for !g.IsGameFinished() {
		g.TakeTurn()
	}
}

func (g *Game[C]) start() {
	iplayers := g.GetPlayers()
	iDeck := g.GetDeck()
	for _, iplayer := range iplayers {
		iplayer.NameSelf()
	}
	iDeck.Shuffle()
}
