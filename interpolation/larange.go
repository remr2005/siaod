package interpolation

import (
	"siaodMath/makefunctions"
)

func Larange1(f func(float64) float64, x []float64) func(float64) float64 {
	arr_funcs := make([]func(float64) float64, len(x))
	for i := 0; i < len(x); i++ {
		arr_funcs[i] = func(x_ float64) float64 { return 1 }
		for j := 0; j < len(x); j++ {
			if j == i {
				continue
			}
			var buf = makefunctions.Addconst(makefunctions.BasicFunc, -x[j])
			buf = makefunctions.Prodconst(buf, 1/(x[i]-x[j]))
			arr_funcs[i] = makefunctions.ProdFunc(arr_funcs[i], buf)
		}
	}
	var buf = func(x_ float64) float64 { return 0 }
	for ind := 0; ind < len(x); ind++ {
		buf = makefunctions.AddFunc(buf, makefunctions.Prodconst(arr_funcs[ind], f(x[ind])))
	}
	return buf
}

func Larange(y, x []float64) func(float64) float64 {
	arr_funcs := make([]func(float64) float64, len(x))
	for i := 0; i < len(x); i++ {
		arr_funcs[i] = func(x_ float64) float64 { return 1 }
		for j := 0; j < len(x); j++ {
			if j == i {
				continue
			}
			var buf = makefunctions.Addconst(makefunctions.BasicFunc, -x[j])
			buf = makefunctions.Prodconst(buf, 1/(x[i]-x[j]))
			arr_funcs[i] = makefunctions.ProdFunc(arr_funcs[i], buf)
		}
	}
	var buf = func(x_ float64) float64 { return 0 }
	for ind := 0; ind < len(x); ind++ {
		buf = makefunctions.AddFunc(buf, makefunctions.Prodconst(arr_funcs[ind], y[ind]))
	}
	return buf
}
