package model

import (
	"fmt"
	"time"
)

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (c *Card) checkBalance() error {
	fmt.Println("[Card] Query to Bank to check balance")
	time.Sleep(time.Second)
	return c.Bank.checkBalance(c.Name)
}
