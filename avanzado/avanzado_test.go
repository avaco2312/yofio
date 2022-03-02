package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"yofio/test_data"
)

const (
	apiUrl    = "https://6xiekuxzq2.execute-api.us-west-2.amazonaws.com/produccion/"
)

func TestAvanzado(t *testing.T) {
	testdata := test_data.CreaTestData()
	for j := 0; j < len(testdata); j++ {
		respuesta, err := http.Post(apiUrl+"credit-assignment", "application/json", strings.NewReader(`{ "investment": `+fmt.Sprint(testdata[j].ImpAsignar)+` }`))
		if err != nil {
			log.Fatal(err)
		}
		if respuesta.StatusCode == http.StatusInternalServerError {
			log.Fatal("Internal server error")
		}
		var resp map[string]interface{}
		err = json.NewDecoder(respuesta.Body).Decode(&resp)
		if err != nil {
			log.Fatal(err)
		}
		respuesta.Body.Close()
		if respuesta.StatusCode == http.StatusOK {
			testdata[j].Id = resp["id"].(string)
			testdata[j].ImpResultado = int32(resp["credit_type_300"].(float64))*300 +
				int32(resp["credit_type_500"].(float64))*500 + int32(resp["credit_type_700"].(float64))*700
			testdata[j].Recibido = true
		} else {
			testdata[j].Id = resp["id"].(string)
			testdata[j].Recibido = false
		}
	}
	var aexcan, aeximp, anoexcan, anoeximp int32
	for j, td := range testdata {
		if td.Esperado != td.Recibido {
			t.Fatalf("Esperado: %v Resultado: %v Asignar: %d Asignado: %d j %d", td.Esperado, td.Recibido, td.ImpAsignar, td.ImpResultado, j)
		} else {
			if td.Esperado && td.ImpAsignar != td.ImpResultado {
				t.Fatalf("Esperado: %v Resultado: %v Asignar: %d Asignado: %d j %d", td.Esperado, td.Recibido, td.ImpAsignar, td.ImpResultado, j)
			}

			if td.Esperado {
				aexcan++
				aeximp += td.ImpAsignar
			} else {
				if td.Id != "0" {
					anoexcan++
					anoeximp += td.ImpAsignar
				}
			}

		}
	}
	promexito := 0.0
	if aexcan != 0 {
		promexito = float64(aeximp) / float64(aexcan)
	}
	promnoexito := 0.0
	if anoexcan != 0 {
		promnoexito = float64(anoeximp) / float64(anoexcan)
	}
	bodyc := fmt.Sprintf(`{"asignaciones_realizadas": %d, "asignaciones_exitosas": %d, "asignaciones_no_exitosas": %d, `+
		`"promedio_inversión_exitosa": %.2f, "promedio_inversión_no_exitosa": %.2f}`,
		aexcan+anoexcan, aexcan, anoexcan, promexito, promnoexito)

	respuesta, err := http.Get(apiUrl + "statistics")
	if err != nil {
		log.Fatal(err)
	}
	if respuesta.StatusCode != http.StatusOK {
		log.Fatal("Internal server error")
	}
	body, err := ioutil.ReadAll(respuesta.Body)
	if err != nil {
		log.Fatal(err)
	}
	respuesta.Body.Close()
	if string(body) != bodyc {
		t.Fatalf("\nCalculado: %s\nBase de Datos: %s", bodyc, body)
	}
}
