package hanlder

import (
	"calculate-service/pkg/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type calculatedResponse struct {
	a int
	b int
}

func Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	a := service.CalculateFactorial(3)
	b := service.CalculateFactorial(3)
	resp := toResponse(a, b)
	fmt.Fprintf(w, resp.String())
}

func toResponse(a int, b int) calculatedResponse {
	return calculatedResponse{a: a, b: b}
}

func (c *calculatedResponse) String() string {
	return fmt.Sprintf("%d %d", c.a, c.b)
}
