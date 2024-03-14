package snau

import "math"

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
