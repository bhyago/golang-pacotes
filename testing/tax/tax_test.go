package tax

import "testing"

func TestCalculateText(t *testing.T) {
	amount := 500.0
	expected := 5.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTextBatch(t *testing.T) {
	type calcTax struct {
		amount   float64
		expected float64
	}

	table := []calcTax{
		{500, 5.0},
		{1000, 10.0},
		{1500, 10.0},
	}

	for _, test := range table {
		result := CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	amount := 500.0
	expected := 5.0

	for i := 0; i < b.N; i++ {
		result := CalculateTax(amount)
		if result != expected {
			b.Errorf("Expected %f, got %f", expected, result)
		}
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	amount := 500.0
	expected := 5.0

	for i := 0; i < b.N; i++ {
		result := CalculateTax2(amount)
		if result != expected {
			b.Errorf("Expected %f, got %f", expected, result)
		}
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1500.0, 1501.0}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Expected positive value, got %f", result)
		}

		if amount > 20000 && result != 20 {
			t.Errorf("Expected 20, got %f", result)
		}
	})
}
