package pkg

type House struct {
	square float64
	town   string
	floors int
}

func NewHouse(square float64, town string, floors int) House {
	return House{
		square: square,
		town:   town,
		floors: floors,
	}
}

func (h *House) GetSquare() float64 {
	return h.square
}

func (h *House) GetTown() string {
	return h.town
}

func (h *House) GetFloors() int {
	return h.floors
}
