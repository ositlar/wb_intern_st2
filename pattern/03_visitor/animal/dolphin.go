package animal

import "ositlar.com/pattern/03_visitor/interfaces"

type dolphin struct {
	roar string
}

func NewDolphin(roar string) *dolphin {
	return &dolphin{
		roar: roar,
	}
}

func (l *dolphin) Accept(v interfaces.Visitor) {
	v.VisitForDolphin(l)
}

func (l *dolphin) GetType() string {
	return "Dolphin"
}
