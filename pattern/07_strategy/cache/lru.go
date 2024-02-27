package cache

import (
	"fmt"

	"ositlar.com/pattern/07_strategy/interfaces"
)

type lru struct {
}

func NewLru() *lru {
	return &lru{}
}

func (l *lru) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by lru strategy")
}
