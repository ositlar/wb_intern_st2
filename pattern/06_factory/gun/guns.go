package gun

import "ositlar.com/pattern/06_factory/interfaces"

type ak47 struct {
	gun
}

func NewAk47() interfaces.Gun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type maverick struct {
	gun
}

func NewMaverick() interfaces.Gun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}
