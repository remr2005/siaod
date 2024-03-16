package interpolation

import (
	"math"

	"gonum.org/v1/gonum/integrate"
)

func sim(f func(float64) float64, a, b float64) float64 {
	x := make([]float64, 0)
	y := make([]float64, 0)
	for i := a; i < b; i += 0.01 {
		x = append(x, i)
		y = append(y, f(i))
	}
	x = append(x, b)
	y = append(y, f(b))
	return integrate.Simpsons(x, y)
}

func Furie(f func(float64) float64, l, N float64, x []float64) []float64 {
	cof := make([]float64, 0)
	cof = append(cof, sim(f, -l, l)/l)
	for i := 1.0; i < N+1; i++ {
		an := func(x float64) float64 {
			return f(x) * math.Cos((math.Pi*i*x)/l)
		}
		bn := func(x float64) float64 {
			return f(x) * math.Sin((math.Pi*i*x)/l)
		}
		cof = append(cof, sim(an, -l, l)/l)
		cof = append(cof, sim(bn, -l, l)/l)
	}
	result := make([]float64, 0)
	for _, val := range x {
		res := cof[0] / 2
		for n := 1.0; n < N+1; n++ {
			res += cof[2*int(n)-1]*math.Cos((math.Pi*n*val)/l) + cof[2*int(n)]*math.Sin((math.Pi*n*val)/l)
		}
		result = append(result, res)
	}
	return result
}

func ComputeMAE(arr1, arr2 []float64) float64 {
	if len(arr1) != len(arr2) {
		panic("Arrays have different lengths")
	}

	var sumError float64
	for i := 0; i < len(arr1); i++ {
		sumError += math.Abs(arr1[i] - arr2[i])
	}

	return sumError / float64(len(arr1))
}
