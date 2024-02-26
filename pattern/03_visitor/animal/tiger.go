package animal

import "ositlar.com/pattern/03_visitor/interfaces"

type tiger struct {
	roar string
}

func NewTiger(roar string) *tiger {
	return &tiger{
		roar: roar,
	}
}

func (l *tiger) Accept(v interfaces.Visitor) {
	v.VisitForTiger(l)
}

func (l *tiger) GetType() string {
	return "Tiger"
}
