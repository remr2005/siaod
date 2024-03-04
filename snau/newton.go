package snau

import (
	"math"
	"siaodMath/derivation"
)

func Newton(infunc func(float64) float64, x, e float64) float64 {
	f := infunc(x)
	f1 := derivation.FirstDeriv(infunc, x)
	for math.Abs(f/f1) > e {
		x -= f / f1
		f = infunc(x)
		f1 = derivation.FirstDeriv(infunc, x)
	}
	return x
}
