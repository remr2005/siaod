package main

import (
	"fmt"
	"siaodMath/integrals"
)

func main() {
	res, pog, err := integrals.Simpson(integrals.F1, 0, 1.6, 2, 1000, 0.0001)
	if err != nil {
		panic(err)
	}
	fmt.Println(res, pog)
	res, pog = integrals.MonteKarlo(integrals.F1, 0, 1.6, 100000, 10)
	fmt.Println(res, pog)
}
