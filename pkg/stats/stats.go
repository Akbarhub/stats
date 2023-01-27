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