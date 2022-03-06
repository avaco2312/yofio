package asigna

import (
	"fmt"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

func NewCreditAssigner() CreditAssigner {
	return NewAsignaCredito()
}

type AsignaCredito struct{}

func NewAsignaCredito() *AsignaCredito {
	return &AsignaCredito{}
}

func (c *AsignaCredito) Assign(investment int32) (int32, int32, int32, error) {
	// Aunque esta garantizado, verificamos que el total es positivo y divisible por 100
	if (investment <= 0) || (investment%100) != 0 {
		return 0, 0, 0, fmt.Errorf("error: %d no es positivo o no es múltiplo de 100", investment)
	}
	// Trabajar con la ecuación simplificada
	f := investment / 100

	// Calculamos el límite superior de v
	vMax := investment / 700

	var v, u int32
	for v = 0; v <= vMax; v++ {
		// Haciendo el calculo de los límites con enteros
		li := (f - 7*v) / 3
		if (f-7*v)%3 != 0 {
			li++
		}
		ls := (2*f - 14*v) / 5
		for u = li; u <= ls; u++ {
			return 2*f - 14*v - 5*u, -f + 7*v + 3*u, v, nil
		}
	}
	return 0, 0, 0, fmt.Errorf("error: %d no es distribuible", investment)
}
