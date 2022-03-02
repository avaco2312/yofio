package asigna

import (
	"sync"
	"testing"
	"yofio/test_data"
)

const nro_gosub = 20 // Número de subrutinas de test concurrentes

func TestDiofanto(t *testing.T) {
	var wg sync.WaitGroup
	ncredito := NewCreditAssigner()
	testdata := test_data.CreaTestData()
	var ch = make(chan struct{}, nro_gosub)
	for j := 0; j < len(testdata); j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			ch <- struct{}{}
			x, y, z, err := ncredito.Assign(testdata[j].ImpAsignar)
			if err == nil {
				testdata[j].ImpResultado = x*300 + y*500 + z*700
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
