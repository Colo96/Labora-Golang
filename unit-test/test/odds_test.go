package test

import "testing"

type Test struct {
	n           int
	expectedSum int
}

func TestSumOddNumbers(t *testing.T) {
	tests := []Test{
		{n: 1, expectedSum: 1},
		{n: 2, expectedSum: 4},
		{n: 3, expectedSum: 9},
		{n: 4, expectedSum: 16},
		{n: 5, expectedSum: 25},
		{n: 10, expectedSum: 100},
	}

	for _, test := range tests {
		result := SumOddNumbers(test.n)
		if result != test.expectedSum {
			t.Errorf("SumOddNumbers(%d) = %d; want %d", test.n, result, test.expectedSum)
		}
	}
}

func BenchmarkSumOddNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOddNumbers(8)
	}
}
