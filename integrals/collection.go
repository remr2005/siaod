package integrals

import "math"

func F1(x float64) float64 {
	return math.Cos(3*x) * math.Pow(math.Cos(x), 3)
}

func F2(x float64) float64 {
	return math.Sin(0.4*x) * math.Pow(math.Sin(x), 0.4)
}

func F3(x float64) float64 {
	return math.Sqrt(math.Sin(x)) * math.Cos(0.5*x)
}

func F4(x float64) float64 {
	return math.Cos(3*x) / (1 + 0.7*math.Cos(x))
}

func F5(x float64) float64 {
	return 1 / math.Pow(0.5*math.Sin(x)+3*math.Cos(x), 2)
}

func F6(x float64) float64 {
	return 1 / (1 - 0.49*math.Pow(math.Sin(x), 2))
}
