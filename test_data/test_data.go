package test_data

import (
	"math/rand"
	"time"
)

const (
	cant_pruebas                 = 200
	porc_no_divisibles_100       = 10
	porc_no_positivos            = 5 + porc_no_divisibles_100
	porc_casos_conocidos_no_dist = 5 + porc_no_positivos
)

var casos_conocidos_no_dist = [...]int32{100, 200, 400}

type TestData struct {
	ImpAsignar,
	ImpResultado int32
	Esperado,
	Recibido bool
	Id string
}

func CreaTestData() []TestData {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	test := []TestData{}
	for i := 0; i <= cant_pruebas; i++ {
		caso := random.Intn(100)
		switch {
		case caso < porc_no_divisibles_100:
			imp := int32(random.Intn(500)*100 + random.Intn(99) + 1)
			test = append(test, TestData{ImpAsignar: imp, ImpResultado: 0, Esperado: false, Recibido: false})
		case caso < porc_no_positivos:
			imp := int32(-random.Intn(500000))
			test = append(test, TestData{ImpAsignar: imp, ImpResultado: 0, Esperado: false, Recibido: false})
		case caso < porc_casos_conocidos_no_dist:
			imp := casos_conocidos_no_dist[random.Intn(len(casos_conocidos_no_dist))]
			test = append(test, TestData{ImpAsignar: imp, ImpResultado: 0, Esperado: false, Recibido: false})
		default:
			imp := int32(random.Intn(500)*300 + random.Intn(500)*500 + random.Intn(500)*700)
			test = append(test, TestData{ImpAsignar: imp, ImpResultado: 0, Esperado: true, Recibido: false})
		}
	}
	return test
}
