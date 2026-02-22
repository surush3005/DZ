package card

import (
	"github.com/surush3005/DZ/pkg/types"
)

// задание 7.1 Зачисление средств
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

// задание 7.2 Процент на остаток
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}

	if card.Balance > 0 {
		return
	}

	if percent > 5_000 {
		return
	}

	bonus := int(card.MinBalance) * percent * daysInMonth / daysInYear
	card.MinBalance += types.Money(bonus)
}
