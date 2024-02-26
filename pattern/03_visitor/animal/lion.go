package animal

import "ositlar.com/pattern/03_visitor/interfaces"

type lion struct {
	roar string
}

func NewLion(roar string) *lion {
	return &lion{
		roar: roar,
	}
}

func (l *lion) Accept(v interfaces.Visitor) {
	v.VisitForLion(l)
}

func (l *lion) GetType() string {
	return "Lion"
}
