package card_game_template

type IPlayer[T any] interface {
	ChooseHand() *T
	NameSelf()
	AddHand(card *T)
	HasFinishDraw() bool
	GetName() string
	ExtractCard(idx int) *T
}
