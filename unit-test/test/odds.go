package test

func SumOddNumbers(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += 2*i + 1
	}
	return sum
}
