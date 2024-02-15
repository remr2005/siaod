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

// стандартное отклонение средних
func SredInt(a []int) int {
	var res1_i int
	var res2_i int
	var len_i int
	for _, i := range a {
		res1_i += i * i
		res2_i += i
		len_i += 1
	}
	var res1 float64 = float64(res1_i)
	var res2 float64 = float64(res2_i)
	var len float64 = float64(len_i)
	return int(math.Sqrt(res1/len - (res2*res2)/(len*len)))
}
