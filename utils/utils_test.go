package utils

import (
	"bytes"
	"io"
	"testing"
)

func TestParseBodySuccess(t *testing.T) {
	var body io.ReadCloser
	body = io.NopCloser(bytes.NewBuffer([]byte("{\"a\": 4, \"b\": 5}")))
	_, err := ParseBody(body)
	if err != nil {
		t.Errorf("ParseBody err: %s", err.Error())
	}
}

func TestParseBodyFail(t *testing.T) {
	var body io.ReadCloser
	body = io.NopCloser(bytes.NewBuffer([]byte{}))

	_, err := ParseBody(body)
	if err == nil {
		t.Errorf("ParseBody err: %s", err.Error())
	}
}
