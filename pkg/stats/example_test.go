package stats

import (
	"fmt"

	"github.com/Akbarhub/types/v2/pkg/bank/types"
)

func ExampleAvg () {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 1000,
			Category: "Auto",
			Status: types.StatusFail,
		},
		{
			ID: 2,
			Amount: 1600,
			Category: "Medical",
			Status: types.StatusOk,
		},
		{
			ID: 3,
			Amount: 800,
			Category: "Food",
			Status: types.StatusFail,
		},
		{
			ID: 4,
			Amount: 200,
			Category: "Auto",
			Status: types.StatusInProgress,
		},
	}
	fmt.Println(Avg(payments))
	// Output:
	// 900
}
	
func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 1000,
			Category: "Medical",
			Status: types.StatusFail,
		},
		{
			ID: 2,
			Amount: 1500,
			Category: "Auto",
			Status: types.StatusInProgress,
		},
		{
			ID: 3,
			Amount: 500,
			Category: "Auto",
			Status: types.StatusOk,
		},
		{
			ID: 4,
			Amount: 300,
			Category: "Food",
			Status: types.StatusFail,
		},
	}
	fmt.Println(TotalInCategory(payments, types.Category("Auto")))
	// Output:
	// 2000
}