package tv

import "fmt"

type tv struct {
	isRunning bool
}

func NewTv() *tv {
	return &tv{}
}

func (t *tv) On() {
	t.isRunning = true
	fmt.Println("Tv is on")
}

func (t *tv) Off() {
	t.isRunning = false
	fmt.Println("Tv is off")
}
