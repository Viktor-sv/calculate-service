package utils

import (
	"encoding/json"
	"io"
)

type RequestData struct {
	A int
	B int
}

func ParseBody(r io.ReadCloser) (*RequestData, error) {
	var data RequestData
	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
