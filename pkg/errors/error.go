package errors

import "encoding/json"

const IncorrectInputError = "Incorrect input"

type responseError struct {
	Error string `json:"error"`
}

func ResponseError(msg string) string {
	e := responseError{
		Error: msg,
	}
	res, _ := json.Marshal(e)
	return string(res)
}
