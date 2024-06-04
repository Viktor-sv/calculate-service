package handler

type FactorialRequest struct {
	A uint64 `json:"a"`
	B uint64 `json:"b"`
}

type factorial struct {
	a uint64
	b uint64
}
