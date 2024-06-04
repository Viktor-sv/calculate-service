package utils

import (
	model "calculate-service/pkg/server/model/factorial"
	"encoding/json"
	"io"
)

func ParseBody(r io.ReadCloser) (*model.Factorial, error) {
	var data model.Factorial
	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
