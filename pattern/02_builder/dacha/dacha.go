package dacha

import "ositlar.com/pattern/02_builder/pkg"

type dachaBuilder struct {
	square float64
	town   string
	floors int
}

func NewDachaBuilder() *dachaBuilder {
	return &dachaBuilder{}
}

func (b *dachaBuilder) SetSquare(ns float64) {
	b.square = ns
}

func (b *dachaBuilder) SetTown(nt string) {
	b.town = nt
}
func (b *dachaBuilder) SetFloors(nf int) {
	b.floors = nf
}

func (b *dachaBuilder) GetHouse() pkg.House {
	return pkg.NewHouse(b.square, b.town, b.floors)
}
