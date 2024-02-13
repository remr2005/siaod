package integrals

import (
	"fmt"
	"math"
)

func Rectangle(inFunc func(float64) float64, a, b, n_min, n_max, e float64) (float64, float64, error) {
	calculate := func(a, b, h float64) float64 {
		res := inFunc(a)
		for a < b {
			a += h
			res += inFunc(a)
		}
		res += inFunc(b)
		return res * h
	}
	var Sh_last float64
	for i := n_min; i < n_max; i *= 2 {
		var Sh float64 = calculate(a, b, math.Abs(b-a)/i)
		if math.Abs(Sh-Sh_last) < e {
			return Sh, Runge(Sh, calculate(a, b, math.Abs(b-a)/(i*2)), "rectangle"), nil
		}
		Sh_last = Sh
	}
	return 0.0, 0, fmt.Errorf("заданная точность не соблюдена в области [n_min,n_max]")
}

func Trapezoid(inFunc func(float64) float64, a, b, n_min, n_max, e float64) (float64, float64, error) {
	calculate := func(a, b, h float64) float64 {
		res := inFunc(a)
		for a < b {
			a += h
			res += 2 * inFunc(a)
		}
		res += inFunc(b)
		return res * h / 2
	}
	var Sh_last float64
	for i := n_min; i < n_max; i *= 2 {
		var Sh float64 = calculate(a, b, math.Abs(b-a)/i)
		if math.Abs(Sh-Sh_last) < e {
			return Sh, Runge(Sh, calculate(a, b, math.Abs(b-a)/(i*2)), "trapezoid"), nil
		}
		Sh_last = Sh
	}
	return 0.0, 0, fmt.Errorf("заданная точность не соблюдена в области [n_min,n_max]")
}
