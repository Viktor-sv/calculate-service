package errors

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestResponseErrorSuccess(t *testing.T) {
	e := responseError{
		Error: IncorrectInputError,
	}
	want, _ := json.Marshal(e)
	errMsg := ResponseError(IncorrectInputError)

	if strings.Compare(string(want), errMsg) != 0 {
		t.Errorf("ResponseError() failed; want %s, got %s", want, errMsg)
	}
}
