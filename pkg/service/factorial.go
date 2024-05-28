package service

func CalculateFactorial(n int) int {
	return factorial(n)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
