package main

import (
	"fmt"
	"yofio/asigna"
)

var creditos = [...]int32{-12000, 300, 400, 500, 510, 1000, 1050, 3000, 4000, 7000, 12000}

func main() {
	ncredito := asigna.NewCreditAssigner()
	for _, monto := range creditos {
		x, y, z, err := ncredito.Assign(monto)
		if err != nil {
			fmt.Printf("%d: %s\n", monto, err.Error())
		} else {
			fmt.Printf("%3d * 300 + %3d * 500 + %3d * 700 = %6d\n", x, y, z, x*300+y*500+z*700)
		}
	}
}
