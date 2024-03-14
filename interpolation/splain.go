package interpolation

import (
	"math"
)

func Splain(x, y, xVal []float64) []float64 {
	n := len(x)
	h := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		h[i] = x[i+1] - x[i]
	}
	alpha := make([]float64, n-1)
	for i := 1; i < n-1; i++ {
		alpha[i] = 3/h[i]*(y[i+1]-y[i]) - 3/h[i-1]*(y[i]-y[i-1])
	}
	l := make([]float64, n)
	mu := make([]float64, n)
	z := make([]float64, n)
	l[0] = 1
	mu[0] = 0
	z[0] = 0
	for i := 1; i < n-1; i++ {
		l[i] = 2*(x[i+1]-x[i-1]) - h[i-1]*mu[i-1]
		mu[i] = h[i] / l[i]
		z[i] = (alpha[i] - h[i-1]*z[i-1]) / l[i]
	}
	l[n-1] = 1
	z[n-1] = 0
	c := make([]float64, n)
	b := make([]float64, n)
	d := make([]float64, n)
	for j := n - 2; j >= 0; j-- {
		c[j] = z[j] - mu[j]*c[j+1]
		b[j] = (y[j+1]-y[j])/h[j] - h[j]*(c[j+1]+2*c[j])/3
		d[j] = (c[j+1] - c[j]) / (3 * h[j])
	}
	interpolatedValue := make([]float64, 0)
	var i int
	for i = 0; i < n-1; i++ {
		for j := 0; j < len(xVal); j++ {
			if x[i] <= xVal[j] && xVal[j] < x[i+1] {
				xDiff := xVal[j] - x[i]
				interpolatedValue = append(interpolatedValue, y[i]+b[i]*xDiff+c[i]*math.Pow(xDiff, 2)+d[i]*math.Pow(xDiff, 3))
			}
		}
	}
	return interpolatedValue
}
