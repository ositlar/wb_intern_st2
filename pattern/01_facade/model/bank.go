package model

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (b *Bank) checkBalance(cardNumber string) error {
	fmt.Printf("[Bank] Geting card's balance(%s)\n", cardNumber)
	time.Sleep(200 * time.Millisecond)
	for _, card := range b.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Bank] Not enouth credits")
		}
	}
	fmt.Println("[Bank] Approved")
	return nil
}
