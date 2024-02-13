package main

import (
	"fmt"
	"siaodMath/integrals"
)

func main() {
	res, pog, err := integrals.Simpson(integrals.F1, 0, 1.6, 2, 200000000, 0.000000000001)
	if err != nil {
		panic(err)
	}
	fmt.Println(res, pog)

	res, pog = integrals.MonteKarlo(integrals.F1, 0, 1.6, 100000, 10)
	fmt.Println(res, pog)

	res, pog = integrals.AdaptSimpson(integrals.F1, 0, 1.6, 200000, 0.000000000001)
	fmt.Println(res, pog)
}
