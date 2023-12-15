package card_game_template

type IDeck[C any] interface {
	Shuffle()
	DrawCards(players []IPlayer[C])
	DrawCard() *C
}
