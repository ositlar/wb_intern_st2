package main

import (
	"ositlar.com/pattern/03_visitor/animal"
	"ositlar.com/pattern/03_visitor/visitor"
)

func main() {
	lion := animal.NewLion("Arrrrr")
	tiger := animal.NewTiger("Wrrrrr")
	dolphin := animal.NewDolphin("Quaa")

	roar := visitor.NewRoar()
	lion.Accept(roar)
	tiger.Accept(roar)
	dolphin.Accept(roar)

	color := visitor.NewColor()
	lion.Accept(color)
	tiger.Accept(color)
	dolphin.Accept(color)

}
