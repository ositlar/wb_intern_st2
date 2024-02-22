package main

import "fmt"

var (
	bank = Bank{
		Name:  "NBer",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "crd-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "crd-2",
		Balance: 240,
		Bank:    &bank,
	}
	user1 = User{
		Name: "user-1",
		Card: &card1,
	}
	user2 = User{
		Name: "user-2",
		Card: &card2,
	}
	prod = Product{
		Name:  "Cheese",
		Price: 185,
	}
	shop = Shop{
		Name:     "6ka",
		Products: []Product{prod},
	}
)

/*
Объект Shop и функция Sell - фасад над всей логикой по оплате товаров
*/
func main() {
	fmt.Println("[Bank] New cards")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]", user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[%s]", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err)
	}
}
