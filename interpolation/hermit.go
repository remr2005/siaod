package interpolation

import (
	"math"
	"siaodMath/derivation"
)

func Hermit(f func(float64) float64, x, xVal []float64) []float64 {
	n := len(x) - 1
	a := make([]float64, 0)
	b := make([]float64, 0)
	c := make([]float64, 0)
	d := make([]float64, 0)
	for i := 0; i < n; i++ {
		h := x[i+1] - x[i]
		b = append(b, derivation.FirstDeriv(f, x[i]))
		a = append(a, f(x[i]))
		c = append(c, (3/(h*h))*(f(x[i+1])-f(x[i]))-(2/h)*derivation.FirstDeriv(f, x[i])-(1/h)*derivation.FirstDeriv(f, x[i+1]))
		d = append(d, (-2/(h*h*h))*(f(x[i+1])-f(x[i]))+(1/(h*h))*(derivation.FirstDeriv(f, x[i+1])+derivation.FirstDeriv(f, x[i])))
	}
	interpolatedValue := make([]float64, 0)
	var i int
	for i = 0; i < n; i++ {
		for j := 0; j < len(xVal); j++ {
			if x[i] <= xVal[j] && xVal[j] < x[i+1] {
				xDiff := xVal[j] - x[i]
				interpolatedValue = append(interpolatedValue, a[i]+b[i]*xDiff+c[i]*math.Pow(xDiff, 2)+d[i]*math.Pow(xDiff, 3))
			}
		}
	}
	return interpolatedValue
}
