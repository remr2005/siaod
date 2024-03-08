package snau

import "math"

func Ditohomia(in_func func(float64) float64, a, b, e float64) float64 {
	x := (a + b) / 2
	if math.Abs(in_func(x)) < e {
		return (a + b) / 2
	} else if (in_func(a) < 0 && in_func(x) < 0) || (in_func(a) > 0 && in_func(x) > 0) {
		return Ditohomia(in_func, x, b, e)
	}
	return Ditohomia(in_func, a, x, e)
}
