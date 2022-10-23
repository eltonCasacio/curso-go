package tax

import "testing"

func FuzzCalculateTax2(f *testing.F) {
	seed := []float64{-1000, -100, -1, 0, 1, 10, 100, 1000, 10000}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Recebido %f, mas esperava 0", result)
		}
	})
}
