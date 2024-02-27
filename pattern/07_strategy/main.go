package main

import "ositlar.com/pattern/07_strategy/cache"

func main() {
	lfu := cache.NewLfu()
	cache1 := cache.InitCache(lfu)
	cache1.Add("a", "1")
	cache1.Add("b", "2")
	cache1.Add("c", "3")
	lru := cache.NewLru()
	cache1.SetEvictionAlgo(lru)
	cache1.Add("d", "4")
	fifo := cache.NewFifo()
	cache1.SetEvictionAlgo(fifo)
	cache1.Add("e", "5")
}
