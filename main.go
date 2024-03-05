package main

import (
	"fmt"
	"math"
	"siaodMath/snau"
)

func main() {
	arr := make([]func([]float64) float64, 0)
	arr = append(arr, f1)
	arr = append(arr, f2)
	// arr = append(arr, F1)
	// arr = append(arr, F2)
	// arr = append(arr, F3)
	// arr = append(arr, F4)
	// arr = append(arr, F5)
	// arr = append(arr, F6)
	fmt.Println("Ньютона -", snau.Newton(func(f float64) float64 { return 2*math.Log(f) - 1/f }, 2, 0.0001))
	fmt.Println("Ньютона Равз -", snau.Rafs(arr, []float64{1, 1}, 0.0001, 10000))
}

// Вывод
//0.19634944003810098 3.811902345625858e-09
//0.19634940866640296 2.2137558286505765e-08
//0.19642765260874218 1.3020805412109851e-05
//0.19597239895846208 0.002043287065238123 (различается при запусках из за рандома)
//0.19634891515714303 5.1453733246886735e-05

func F1(x []float64) float64 {
	return 2*x[0] + 1*x[1] + 3*x[2] + 4*x[3] + 5*x[4] + 6*x[5] - 20
}
func F2(x []float64) float64 {
	return 0*x[0] + (-1)*x[1] + 2*x[2] + 3*x[3] + 1*x[4] + 2*x[5] + 3
}
func F3(x []float64) float64 {
	return 1*x[0] + 0*x[1] + 1*x[2] + 2*x[3] + (-1)*x[4] + 1*x[5] - 12
}
func F4(x []float64) float64 {
	return (-1)*x[0] + 2*x[1] + 0*x[2] + 1*x[3] + 3*x[4] + (-2)*x[5] - 5
}
func F5(x []float64) float64 {
	return 4*x[0] + 2*x[1] + 1*x[2] + 0*x[3] + (-1)*x[4] + 3*x[5] - 8
}
func F6(x []float64) float64 {
	return 3*x[0] + 1*x[1] + 2*x[2] + (-1)*x[3] + 0*x[4] + 4*x[5] - 10
}

func f1(x []float64) float64 {
	return math.Sin(x[0]+1) - x[1] - 1.2
}
func f2(x []float64) float64 {
	return 2*x[0] + math.Cos(x[1]) - 2
}
