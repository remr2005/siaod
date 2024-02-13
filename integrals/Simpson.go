package integrals

import (
	"fmt"
	"math"
)

func Simpson(inFunc func(float64) float64, a, b, n_min, n_max, e float64) (float64, float64, error) {
	calculate := func(a, b, h float64) float64 {
		res := inFunc(a)
		var i int = 0
		for a < b {
			a += h
			if i%2 == 0 {
				res += inFunc(a) * 4
			} else {
				res += inFunc(a) * 2
			}
			i += 1
		}
		res += inFunc(b)
		return res * h / 3
	}
	var Sh_last float64
	for i := n_min; i < n_max; i *= 2 {
		var Sh float64 = calculate(a, b, math.Abs(b-a)/i)
		if math.Abs(Sh-Sh_last) < e {
			return Sh, Runge(Sh, calculate(a, b, math.Abs(b-a)/(i*2)), "simpson"), nil
		}
		Sh_last = Sh
	}
	return 0.0, 0, fmt.Errorf("заданная точность не соблюдена в области [n_min,n_max]")
}

func AdaptSimpson(inFunc func(float64) float64, a, b, k, e float64) (float64, float64) {
	calculate := func(a, b, h float64) float64 {
		res := inFunc(a)
		var i int = 0
		for a < b {
			a += h
			if i%2 == 0 {
				res += inFunc(a) * 4
			} else {
				res += inFunc(a) * 2
			}
			i += 1
		}
		res += inFunc(b)
		return res * h / 3
	}
	x := calculate(a, b, (a+b)/k)
	y := calculate(a, b, (a+b)/(k*2))
	if !(math.Abs(x-y) < 15*e) {
		x, _ = AdaptSimpson(inFunc, a, (a+b)/2, k, e/2)
		y, _ = AdaptSimpson(inFunc, (a+b)/2, b, k, e/2)
		return x + y, 0.0
	} else {
		return y, Runge(x, y, "simson")
	}
}
