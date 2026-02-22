package card

import (
	"github.com/surush3005/DZ/pkg/types"
)

func Deposit(card *types.Card, amount types.Money) {
	if card == nil {
		return
	}

	if !card.Active {
		return
	}

	if amount <= 0 {
		return
	}

	const MaxDeposit types.Money = 50_000
	if amount > MaxDeposit {
		return
	}
	card.Balance += amount
}

/*
type Service struct {
	cards    []*types.Card
	payments []*types.Payment
}
/*
func Issue(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		Id:       1000,
		PAN:      "5058 XXXX XXXX 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	/*
		return types.Card{
			Id:       1001,
			PAN:      "5058 XXXX XXXX 9999",
			Balance:  -100,
			Currency: currency,
			Color:    color,
			Name:     name,
			Active:   false,
		}

}
*/
