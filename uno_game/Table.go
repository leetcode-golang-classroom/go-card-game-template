package uno_game

type Table struct {
	cards []*UnoCard
}

func NewTable() *Table {
	return &Table{
		cards: []*UnoCard{},
	}
}

func (t *Table) PushBackToDeck() []*UnoCard {
	remain := t.cards[len(t.cards)-1]
	others := t.cards[:len(t.cards)-1]
	t.cards = append([]*UnoCard{}, remain)
	return others
}

func (t *Table) AddCard(card *UnoCard) {
	t.cards = append(t.cards, card)
}

func (t *Table) FindTopest() *UnoCard {
	return t.cards[len(t.cards)-1]
}
