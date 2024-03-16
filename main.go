package main

import (
	"fmt"
	"math"
	"siaodMath/interpolation"
)

func main() {
	// Lab 4
	y := make([]float64, 0)
	for _, i := range []float64{1.00, 1.05, 1.09, 1.13, 1.15, 1.17, 1.2} {
		y = append(y, math.Exp(i))
	}
	f := interpolation.Larange([]float64{1.1, 1.4, 1.6, 1.7, 1.9}, []float64{1, 2, 3, 4, 5})
	fmt.Println("Лерандж", f(1))
	fmt.Println("Сплайн", interpolation.Splain([]float64{1.00, 1.05, 1.09, 1.13, 1.15, 1.17, 1.2}, y, []float64{1.05, 1.09, 1.13, 1.15, 1.17}))
	fmt.Println("Эрмит", interpolation.Hermit(math.Exp, []float64{1.00, 1.05, 1.09, 1.13, 1.15, 1.17, 1.2}, []float64{1.05, 1.09, 1.13, 1.15, 1.17}))
	// Lab 5
	fmt.Println(interpolation.FurieSet(f, 100, []float64{0, math.Pi / 2, math.Pi}))
	fmt.Println(interpolation.Furie2(f, -2*math.Pi, 2*math.Pi, 1000, []float64{0, math.Pi / 2, math.Pi}))
	fmt.Println(f(0))
}

func f(x float64) float64 {
	return (2 / math.Pi) * math.Asin(math.Sin(x))
}
