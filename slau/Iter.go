package slau

import (
	_ "gonum.org/v1/gonum/mat"
)

// Как писать код?
// сделать бранч
// написать тесты
// написать тестируемые функции
// сохранить код в бранче
// скомпить код в .dll файл
// dll добавить в main бранч
// и так каждую функцию

func Diagonal(a [][]float64) [][]float64 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if i != j {
				a[i][j] = 0
			}
		}
	}
	return a
}

func UpperTriangle(a [][]float64) [][]float64 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if i > j {
				a[i][j] = 0
			}
		}
	}
	return a
}

func DownTriangle(a [][]float64) [][]float64 {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if i < j {
				a[i][j] = 0
			}
		}
	}
	return a
}

/*
func GausZeidel(a [][]float64, b []float64, e float64) {
	x := make([]float64, 0)
	for i := 0; i < len(b); i++ {
		x = append(x, 1)
	}
	X := mat.NewDense(len(b), 1, x)
	D := Array_to_matrix(Diagonal(a))
	L := Array_to_matrix(DownTriangle(a))
	M := mat.NewDense(len(a), len(a[0]), nil)
	M.Add(D, L)
	err := M.Inverse(M)
	r, c := M.Dims()
	C := mat.NewDense(r, c, nil)
	C.Product(M, N)
	if err != nil {
		panic(err)
	}

	B := mat.NewDense(1, len(b), b)

}
*/

func GausZeidel(a [][]float64, b []float64, e float64) []float64 {
	x := make([]float64, 0)
	for i := 0; i < len(b); i++ {
		x = append(x, 1.1)
	}
	x_ := make([]float64, len(x))
	copy(x_, x)

	var calc_c = func(i, j int) float64 {
		if i == j {
			return 0
		} else {
			return -a[i][j] / a[i][i]
		}
	}
	var calc_d = func(i int) float64 { return b[i] / a[i][i] }
	var sub_vect = func(a, b []float64) (res float64) {
		for i := 0; i < len(a); i++ {
			res += a[i] - b[i]
		}
		return res
	}
	for sub_vect(x, x_) <= e {
		copy(x_, x)
		for i := 0; i < len(a); i++ {
			d := calc_d(i)
			var res float64
			for j := 0; j < len(a[0]); j++ {
				res += calc_c(i, j) * x[j]
			}
			x[i] = res + d
		}
	}
	return x
}
