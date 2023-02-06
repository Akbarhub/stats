package stats

import (

	"github.com/Akbarhub/types/v2/pkg/bank/types"
)

// CategoriesAvg возвращает среднюю сумму платежа по каждой категории.
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money{
	categories := map[types.Category]types.Money{}
	count := map[types.Category]int{}
	for _, payment := range payments {
		count[payment.Category]++;
		categories[payment.Category] += payment.Amount
	}
	for category, money := range categories {
		categories[category] = money / types.Money(count[category])
	}
 	return categories
}

