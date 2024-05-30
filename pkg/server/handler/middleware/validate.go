package middleware

import (
	"bytes"
	"calculate-service/pkg/errors"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func Validate(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		bodyBytes, err := io.ReadAll(r.Body)
		r.Body.Close() //  must close
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var b []byte
		_ = json.Unmarshal(bodyBytes, &b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, errors.ResponseError(errors.IncorrectInputError))
			return
		}
		h(w, r, ps)
	}
}
