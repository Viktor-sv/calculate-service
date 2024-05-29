package hanlder

import (
	"calculate-service/pkg/service"
	"calculate-service/utils"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	data, err := utils.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := make(chan int)
	go service.CalculateFactorial(data.A, res)
	go service.CalculateFactorial(data.B, res)

	resp := toResponse(<-res, <-res)
	fmt.Fprintf(w, resp)
}

func toResponse(a int, b int) string {
	e := responseError{
		Error: msg,
	}
	res, _ := json.Marshal(e)
	return string(res)
}
