package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yofio/asigna"
)

type Request struct {
	Investment int32
}

type AsignaMonto struct{}

func (a AsignaMonto) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := Request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Investment <= 0 {
		http.Error(w, "petición incorrecta", http.StatusBadRequest)
		return
	}
	ncredito := asigna.NewCreditAssigner()
	x, y, z, err := ncredito.Assign(req.Investment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		fmt.Fprintf(w, `{"credit_type_300": %d, "credit_type_500": %d, "credit_type_700": %d}`, x, y, z)
	}
}
