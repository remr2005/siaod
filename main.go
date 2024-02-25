package main

import (
	"fmt"
	"siaodMath/slau"

	"gonum.org/v1/gonum/mat"
)

func main() {
	// Лаба 2
	a := [][]float64{{5, 0, 1}, {1, 3, -1}, {-3, 2, 10}}
	b := []float64{11, 4, 6}
	aCopy := make([][]float64, len(a))
	copy(aCopy, a)
	bCopy := make([]float64, len(b))
	copy(bCopy, b)
	bCopyCopy := make([]float64, len(b))
	copy(bCopyCopy, b)

	A := mat.NewDense(3, 3, slau.Two2one(a))
	arr, err := slau.Gaus(a, b)
	if err != nil {
		panic(err)
	}
	arr_ := slau.GausZeidel(aCopy, bCopy, 0.0001)
	X := mat.NewDense(3, 1, arr)
	B := mat.NewDense(3, 1, nil)
	B_ := mat.NewDense(3, 1, nil)
	X_ := mat.NewDense(3, 1, arr_)
	B_.Product(A, X_)
	B.Product(A, X)
	fmt.Println(B.RawMatrix().Data, b)
	fmt.Println(B.RawMatrix().Data, "-", bCopyCopy, "-", B_.RawMatrix().Data)

	// Лаба 3

}

// Вывод
//0.19634944003810098 3.811902345625858e-09
//0.19634940866640296 2.2137558286505765e-08
//0.19642765260874218 1.3020805412109851e-05
//0.19597239895846208 0.002043287065238123 (различается при запусках из за рандома)
//0.19634891515714303 5.1453733246886735e-05
