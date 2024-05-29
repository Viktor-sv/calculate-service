package service

import "sync"

func CalculateFactorial(wg *sync.WaitGroup, n int, res chan int) {
	defer wg.Done()
	res <- factorial(n)

	return
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
