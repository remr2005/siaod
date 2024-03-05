package snau

import (
	"math"
	"siaodMath/derivation"
	"siaodMath/slau"

	"gonum.org/v1/gonum/mat"
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

func Rafs(arr_func []func([]float64) float64, x []float64, e float64, bord int) []float64 {
	W := make([][]float64, 0)
	for i := 0; i < len(arr_func); i++ {
		arr_ := make([]float64, len(x))
		W = append(W, arr_)
	}
	var calc_W = func(a_ [][]float64) [][]float64 {
		a := make([][]float64, len(a_))
		copy(a, a_)
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a[i]); j++ {
				a[i][j] = derivation.FirstDerivArray(arr_func[i], x, j)
			}
		}
		return a
	}
	W = calc_W(W)

	F := make([]float64, 0)
	for i := 0; i < len(arr_func); i++ {
		F = append(F, arr_func[i](x))
	}
	var calc_del_X = func(f []float64, W [][]float64) []float64 {
		F = slau.Scale(-1, f)
		matrixF := mat.NewDense(len(W), 1, F)
		matrix := mat.NewDense(len(W), len(W[0]), slau.Two2one(W))

		matrix.Inverse(matrix)
		mat := mat.NewDense(len(x), 1, nil)
		mat.Product(matrix, matrixF)
		delta_x := mat.RawMatrix().Data
		return delta_x
	}
	del_x := calc_del_X(F, W)
	x = slau.AddRow(x, del_x)
	sigma := slau.MinArray(x)
	k := 0
	for sigma > e {
		k += 1
		if k > bord {
			break
		}
		W = calc_W(W)
		for i := 0; i < len(arr_func); i++ {
			F[i] = arr_func[i](x)
		}
		del_x = calc_del_X(F, W)
		x = slau.AddRow(x, del_x)
		sigma = slau.MinArray(x)
	}
	return x
}
