package main

import (
	"fmt"

	"github.com/surush3005/DZ.git/pkg/types"
	"github.com/surush3005/DZ.git/pkg/card"
)

func main() {

	cards := &types.Card{
		Id:       0001,
		PAN:      "5058 XXXX XXXX 9999",
		Balance:  -100,
		Currency: types.TJS,
		Color:    "Black",
		Name:     "MyCard",
		Active:   true,
	}

	card.Deposit(card,5000)

	fmt.Println("Новый баланс:", cards.Balance)
}
