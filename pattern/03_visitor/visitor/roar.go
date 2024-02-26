package visitor

import (
	"fmt"

	"ositlar.com/pattern/03_visitor/interfaces"
)

type roar struct {
	roar string
}

func NewRoar() *roar {
	return &roar{}
}

func (r *roar) VisitForLion(a interfaces.Animal) {
	r.roar = "Lion's roar"
	fmt.Println(r.roar)
}
func (r *roar) VisitForTiger(a interfaces.Animal) {
	r.roar = "Tiger's roar"
	fmt.Println(r.roar)
}
func (r *roar) VisitForDolphin(a interfaces.Animal) {
	r.roar = "Dolphin's roar"
	fmt.Println(r.roar)
}
