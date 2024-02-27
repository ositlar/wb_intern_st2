package gun

import (
	"errors"

	"ositlar.com/pattern/06_factory/interfaces"
)

func GetGun(gunType string) (interfaces.Gun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "maverick" {
		return NewMaverick(), nil
	}
	return nil, errors.New("wrong gun type passed")
}
