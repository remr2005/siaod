package main

import (
	"fmt"
	"math"
	"siaodMath/interpolation"
)

func main() {
	y := make([]float64, 0)
	for _, i := range []float64{1.00, 1.05, 1.09, 1.13, 1.15, 1.17, 1.2} {
		y = append(y, math.Exp(i))
	}
	f := interpolation.Larange([]float64{1.1, 1.4, 1.6, 1.7, 1.9}, []float64{1, 2, 3, 4, 5})
	fmt.Println(f(1))
	fmt.Println(interpolation.Splain([]float64{1.00, 1.05, 1.09, 1.13, 1.15, 1.17, 1.2}, y, []float64{1.05, 1.09, 1.13, 1.15, 1.17}))
	fmt.Println(interpolation.Hermit(math.Exp, []float64{1.00, 1.05, 1.09, 1.13, 1.15, 1.17, 1.2}, []float64{1.05, 1.09, 1.13, 1.15, 1.17}))
}
