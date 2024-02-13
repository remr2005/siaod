package integrals

import (
	"math"
	"math/rand"
)

func MonteKarlo(inFunc func(x float64) float64, a, b, n float64, iter int) (float64, float64) {
	var res float64
	var u float64 = rand.Float64()*(b-a) + a
	arr := make([]float64, 0)
	for j := 0; j < iter; j++ {
		for i := 1.0; i < n; i++ {
			res += inFunc(u)
			u = rand.Float64()*math.Abs(a-b) + a
		}
		arr = append(arr, (res*(b-a))/n)
		res = 0
	}
	return arr[0], Sred(arr)
}
