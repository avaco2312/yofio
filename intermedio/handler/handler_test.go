package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"yofio/test_data"
)

const nro_gosub = 20 // Número de subrutinas de test concurrentes

func TestIntermedio(t *testing.T) {
	handler := &AsignaMonto{}
	var wg sync.WaitGroup
	testdata := test_data.CreaTestData()
	var ch = make(chan struct{}, nro_gosub)
	for j := 0; j < len(testdata); j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			ch <- struct{}{}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/credit-assignment", bytes.NewBuffer([]byte(`{ "investment": `+fmt.Sprint(testdata[j].ImpAsignar)+` }`)))
			handler.ServeHTTP(w, r)
			if w.Result().StatusCode == http.StatusInternalServerError {
				log.Fatal("Internal server error")
			}
			if w.Result().StatusCode == http.StatusOK {
				var resp map[string]int32
				err := json.NewDecoder(w.Body).Decode(&resp)
				if err != nil {
					log.Fatal("Error en respuesta del servidor")
				}
				testdata[j].ImpResultado = resp["credit_type_300"]*300 + resp["credit_type_500"]*500 + resp["credit_type_700"]*700
				testdata[j].Recibido = true
			} else {
				testdata[j].Recibido = false
			}
			<-ch
		}(j)
	}
	wg.Wait()
	for j, td := range testdata {
		if td.Esperado != td.Recibido {
			t.Fatalf("Esperado: %v Resultado: %v Asignar: %d Asignado: %d j %d", td.Esperado, td.Recibido, td.ImpAsignar, td.ImpResultado, j)
		} else {
			if td.Esperado && td.ImpAsignar != td.ImpResultado {
				t.Fatalf("Esperado: %v Resultado: %v Asignar: %d Asignado: %d j %d", td.Esperado, td.Recibido, td.ImpAsignar, td.ImpResultado, j)
			}
		}
	}
}
