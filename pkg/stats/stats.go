package stats

import(
	"github.com/Akbarhub/types/pkg/bank/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money{
	AverageSum := int64(0)
	Aveg := len(payments)
	for _, payment := range payments {
		AverageSum += int64(payment.Amount)
		
	}
	AverageSum = AverageSum / int64(Aveg)
		return types.Money(AverageSum)	
	
}

// TotalInCategory находит сумму пакупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	
	totall := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category{
			totall += payment.Amount
		}
	}
	return totall
}