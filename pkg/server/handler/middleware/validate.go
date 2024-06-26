package middleware

import (
	"bytes"
	"calculate-service/pkg/errors"
	"calculate-service/pkg/server/handler"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func Validate(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body.Close() //  must close
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		var b handler.FactorialRequest
		err := json.Unmarshal(bodyBytes, &b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", errors.ResponseError(errors.IncorrectInputError))
			return
		}
		h(w, r, ps)
	}
}
