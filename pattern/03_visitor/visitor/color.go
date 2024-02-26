package visitor

import (
	"fmt"

	"ositlar.com/pattern/03_visitor/interfaces"
)

type color struct {
	color string
}

func NewColor() *color {
	return &color{}
}

func (c *color) VisitForLion(s interfaces.Animal) {
	c.color = "Brown"
	fmt.Println(c.color)
}

func (c *color) VisitForTiger(s interfaces.Animal) {
	c.color = "Orange"
	fmt.Println(c.color)
}

func (c *color) VisitForDolphin(s interfaces.Animal) {
	c.color = "Grey"
	fmt.Println(c.color)
}
