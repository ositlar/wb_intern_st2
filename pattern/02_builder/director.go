package main

import (
	"ositlar.com/pattern/02_builder/model"
	"ositlar.com/pattern/02_builder/pkg"
)

type director struct {
	builder model.Builder
}

func NewDirector(b model.Builder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) SetBuilder(b model.Builder) {
	d.builder = b
}

func (d *director) BuildHouse(s float64, t string, f int) pkg.House {
	d.builder.SetSquare(s)
	d.builder.SetTown(t)
	d.builder.SetFloors(f)
	return d.builder.GetHouse()
}
