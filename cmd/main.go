package main

import (
	"fmt"

	"github.com/surush3005/DZ.git/pkg/card"
	"github.com/surush3005/DZ/pkg/types"
)

func main() {
	//задание 7.1 Зачисление средств
	crd := &types.Card{
		Id:       0001,
		PAN:      "5058 XXXX XXXX 9999",
		Balance:  -100,
		Currency: types.TJS,
		Color:    "Black",
		Name:     "MyCard",
		Active:   true,
	}

	card.Deposit(crd, 5000)

	fmt.Println("Новый баланс:", crd.Balance)

	//задание 7.2 Процент на остаток
	bonus := &types.Card{
		Id:         0002,
		PAN:        "5058 XXXX XXXX 9999",
		Balance:    100,
		MinBalance: 3000,
		Currency:   types.TJS,
		Color:      "Black",
		Name:       "MyCard",
		Active:     true,
	}

	card.AddBonus(bonus, 3, 30, 365)

	fmt.Println("Доход:", MinBalance)
}
