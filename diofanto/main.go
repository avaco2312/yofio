package main

import (
	"fmt"
	"math"
)

var creditos = [...]int32{-12000, 300, 400, 500, 510, 1000, 1050, 3000, 4000, 7000, 12000}

func main() {
	for _, investment := range creditos {
		// Aunque esta garantizado, verificamos que el total es positivo y divisible por 100
		if (investment <= 0) || (investment%100) != 0 {
			fmt.Printf("%d\terror: no es positivo o no es múltiplo de 100\n", investment)
			continue
		}
		// Trabajar con la ecuación simplificada
		f := int32(investment / 100)

		// Calculamos el límite superior de v
		vMax := int32(investment / 700)

		var v, u int32
		fmt.Print(investment)
		found := false
		for v = 0; v <= vMax; v++ {
			for u = int32(math.Ceil((float64(f - 7*v)) / 3.0)); u <= (2*f-14*v)/5; u++ {
				x := 2*f - 14*v - 5*u
				y := -f + 7*v + 3*u
				fmt.Printf("\t%3d * 300 + %3d * 500 + %3d * 700 = %6d\n", x, y, v, x*300+y*500+v*700)
				found = true
			}
			if !found {
				fmt.Printf("\tno tiene distribución válida\n")
			}
		}
	}
}
