package service

func CalculateFactorial(n int, res chan int) {
	res <- factorial(n)
	return
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
