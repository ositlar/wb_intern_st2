package model

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse()
}
