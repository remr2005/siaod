package integrals

import "math"

func Runge(a float64, b float64, metod string) float64 {
	var k float64 = 4
	if metod == "rectangle" || metod == "trapezoid" {
		k = 2
	}
	return math.Abs(a-b) / (math.Pow(2, k) - 1)
}

func Sred(a []float64) float64 {
	var res1 float64
	var res2 float64
	var len float64
	for _, i := range a {
		res1 += i * i
		res2 += i
		len += 1
	}
	return math.Sqrt(res1/len - (res2*res2)/(len*len))
}
