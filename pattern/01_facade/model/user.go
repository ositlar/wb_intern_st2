package model

type User struct {
	Name string
	Card *Card
}

func (u *User) getBalance() float64 {
	return u.Card.Balance
}
