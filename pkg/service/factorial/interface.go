package factorial

import (
	"calculate-service/pkg/server/model/factorial"
	"context"
)

type Service interface {
	Factorial(ctx context.Context, f *model.Factorial) (*model.Factorial, error)
}
