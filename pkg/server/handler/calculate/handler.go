package hanlder

import (
	"calculate-service/pkg/service"
	"calculate-service/utils"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sync"
)

func Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	data, err := utils.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var wg sync.WaitGroup
	a := make(chan int)
	b := make(chan int)
	wg.Add(2)
	go service.CalculateFactorial(&wg, data.A, a)
	go service.CalculateFactorial(&wg, data.B, b)

	resp := toResponse(<-a, <-b)
	wg.Wait()
	fmt.Fprintf(w, resp)
}

func toResponse(a int, b int) string {
	resp := FactorialResp{
		A: a,
		B: b,
	}
	res, _ := json.Marshal(resp)
	return string(res)
}
