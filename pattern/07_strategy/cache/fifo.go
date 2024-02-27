package cache

import (
	"fmt"

	"ositlar.com/pattern/07_strategy/interfaces"
)

type fifo struct {
}

func NewFifo() *fifo {
	return &fifo{}
}

func (l *fifo) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by fifo strategy")
}
