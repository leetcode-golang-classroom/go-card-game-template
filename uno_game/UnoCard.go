package uno_game

import "fmt"

type UnoCard struct {
	color  Color
	number Number
}

var colors = []string{
	"BLUE",
	"RED",
	"YELLOW",
	"GREEN",
}

func (c *UnoCard) String() string {
	return fmt.Sprintf("[%v%v]", colors[c.color], c.number)
}
func (c *UnoCard) Equal(card *UnoCard) bool {
	return c.color == card.color || c.number == card.number
}

type Color int

const (
	BLUE Color = iota
	RED
	YELLOW
	GREEN
)

type Number int

const (
	Zero Number = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

var Colors = []Color{
	BLUE,
	RED,
	YELLOW,
	GREEN,
}
var Numbers = []Number{
	Zero,
	One,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
}
