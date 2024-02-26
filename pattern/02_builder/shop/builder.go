package shop

import "ositlar.com/pattern/02_builder/pkg"

type shopBuilder struct {
	square float64
	town   string
	floors int
}

func NewShopBuilder() *shopBuilder {
	return &shopBuilder{}
}

func (b *shopBuilder) SetSquare(ns float64) {
	b.square = ns
}

func (b *shopBuilder) SetTown(nt string) {
	b.town = nt
}
func (b *shopBuilder) SetFloors(nf int) {
	b.floors = nf
}

func (b *shopBuilder) GetHouse() pkg.House {
	return pkg.NewHouse(b.square, b.town, b.floors)
}
