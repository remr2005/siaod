package main

import (
	"fmt"
	"siaodMath/integrals"
)

func main() {

	// Лаба 1
	// res, pog, err := integrals.Simpson(integrals.F1, 0, 1.6, 10, 1000, 0.0001)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res, pog)

	// res, pog, err = integrals.Trapezoid(integrals.F1, 0, 1.6, 10, 1000, 0.0001)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res, pog)

	// res, pog, err = integrals.Rectangle(integrals.F1, 0, 1.6, 10, 100000, 0.0001)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res, pog)

	// res, pog = integrals.MonteKarlo(integrals.F1, 0, 1.6, 100000, 10)
	// fmt.Println(res, pog)

	// res, pog = integrals.AdaptSimpson(integrals.F1, 0, 1.6, 10, 0.0001)
	// fmt.Println(res, pog)

	fmt.Println(integrals.MonteKarloImage2())
	// Лаба 2
	// a := make([][]float64, 0)
	// a = append(a, []float64{2, 0, -1})
	// a = append(a, []float64{-1, 3, -1})
	// a = append(a, []float64{1, -1, 4})
	// b := make([]float64, 0)
	// b = append(b, -3)
	// b = append(b, 2)
	// b = append(b, 3)
	// slau.Print_Arr(a)
	// vect, err := slau.Gaus(a, b)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(vect)
	// fmt.Println(slau.Two2one(a))

	// Лаба 3

}

// Вывод
//0.19634944003810098 3.811902345625858e-09
//0.19634940866640296 2.2137558286505765e-08
//0.19642765260874218 1.3020805412109851e-05
//0.19597239895846208 0.002043287065238123 (различается при запусках из за рандома)
//0.19634891515714303 5.1453733246886735e-05
