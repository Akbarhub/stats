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

// PeriodsDynamic сравнивает расходы по категориям за два периода.
func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money,
)map[types.Category]types.Money{
	prDinamic := map[types.Category]types.Money{}
	for key, value := range second {
		_, status := first[key]
		if	!status{
			prDinamic[key] = value
		}
		prDinamic[key] = value - first[key]
	}
	for key := range first{
		_, status := second[key]
		if	!status {
			prDinamic[key] = 0
		}
	}
	
	return prDinamic
}