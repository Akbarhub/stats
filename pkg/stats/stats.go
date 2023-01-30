package stats

import (
	"github.com/Akbarhub/types/v2/pkg/bank/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money{
	avaregeSum := types.Money(0)
	
	i := types.Money(0)

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			i++
			avaregeSum += payment.Amount
		}
	}
		avaregeSum = avaregeSum / i
	return types.Money(avaregeSum)
}

// TotalInCategory находит сумму пакупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	
	totall := types.Money(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category{
			totall += payment.Amount
		}
	}
		return totall
}