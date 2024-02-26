package model

import "ositlar.com/pattern/02_builder/pkg"

type Builder interface {
	SetSquare(float64)
	SetTown(string)
	SetFloors(int)
	GetHouse() pkg.House
}
