package factorial

import (
	model "calculate-service/pkg/server/model/factorial"
	"context"
)

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Factorial(ctx context.Context, data *model.Factorial) (*model.Factorial, error) {
	a := make(chan uint64)
	b := make(chan uint64)

	go CalculateFactorial(data.A, a)
	go CalculateFactorial(data.B, b)

	res := &model.Factorial{
		A: <-a,
		B: <-b,
	}
	return res, nil
}

func CalculateFactorial(n uint64, res chan uint64) {
	res <- factorial(n)
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
