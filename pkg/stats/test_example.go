package stats

import(
	"github.com/Akbarhub/types/pkg/bank/types"
	"fmt"
)

func ExampleAvg () {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 1000,
			Category: "Auto",
		},
		{
			ID: 2,
			Amount: 1500,
			Category: "Medical",
		},
		{
			ID: 3,
			Amount: 500,
			Category: "Food",
		},
		{
			ID: 4,
			Amount: 200,
			Category: "Auto",
		},
	}
	fmt.Println(Avg(payments))
	// Output:
	// 800
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 1000,
			Category: "Auto",
		},
		{
			ID: 2,
			Amount: 1500,
			Category: "Medical",
		},
		{
			ID: 3,
			Amount: 500,
			Category: "Food",
		},
		{
			ID: 4,
			Amount: 300,
			Category: "Auto",
		},
	}
	fmt.Println(TotalInCategory(payments, types.Category("Auto")))
	// Output:
	// 1300
}