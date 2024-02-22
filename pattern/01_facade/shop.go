package main

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (s *Shop) Sell(u User, p string) error {
	fmt.Println("[Shop] Query to User for get balance on card")
	time.Sleep(1000 * time.Millisecond)
	err := u.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Shop] Check - can [%s] buy [%s]\n", u.Name, p)
	time.Sleep(200 * time.Millisecond)
	for _, product := range s.Products {
		if product.Name != p {
			continue
		}
		if product.Price > u.GetBalance() {
			return errors.New("[Shop] Not enouth money on balance")
		}
		fmt.Printf("[Shop] Product [%s] bought by [%s]\n", product.Name, u.Name)
		break
	}
	return nil
}
