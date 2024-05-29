package middleware

import (
	"calculate-service/pkg/errors"
	"calculate-service/utils"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Validate(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_, err := utils.ParseBody(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, errors.ResponseError(errors.IncorrectInputError))
			return
		}
		h(w, r, ps)
	}
}
