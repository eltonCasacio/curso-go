package tax

import "testing"

type calcTax struct {
	amount, expected float64
}

func TestCalculateTax(t *testing.T) {

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
		{0.0, 0.0},
	}

	for _, item := range table {
		result := CalculateTax(item.amount)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}
