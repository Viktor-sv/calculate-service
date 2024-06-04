package hanlder

import (
	model "calculate-service/pkg/server/model/factorial"
	"calculate-service/pkg/service/factorial"
	"calculate-service/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handler struct {
	factorial.Service
}

func NewHandler(s *factorial.Service) *Handler {
	return &Handler{*s}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	data, err := utils.ParseBody(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.Factorial(context.Background(), data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Fprintf(w, toResponse(res))
}

func toResponse(f *model.Factorial) string {
	resp := FactorialResp{
		A: f.A,
		B: f.B,
	}
	res, _ := json.Marshal(resp)
	return string(res)
}
