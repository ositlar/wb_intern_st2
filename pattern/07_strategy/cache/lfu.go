package cache

import (
	"fmt"

	"ositlar.com/pattern/07_strategy/interfaces"
)

type lfu struct {
}

func NewLfu() *lfu {
	return &lfu{}
}

func (l *lfu) Evict(c interfaces.Cache) {
	fmt.Println("Evicting by lfu strategy")
}
