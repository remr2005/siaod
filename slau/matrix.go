package slau

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// replace rows a and b
func Rep_r(a, b int, arr [][]float64) [][]float64 {
	buf := arr[a]
	arr[a] = arr[b]
	arr[b] = buf
	return arr
}

// делает одномерный массив из двумерного
func Two2one(arr [][]float64) (res []float64) {
	res = make([]float64, 0)
	for _, i := range arr {
		for _, j := range i {
			res = append(res, j)
		}
	}
	return res
}

// слайс в матрицу гонам
func Array_to_matrix(arr [][]float64) mat.Dense {
	return *mat.NewDense(len(arr), len(arr[0]), Two2one(arr))
}

// минусовать строки слайса
func SubRows(a, b []float64) (res []float64) {
	var nul int = -1
	for n, i := range a {
		if i == 0 {
			nul = n
		}
	}
	k := b[nul+1] / a[nul+1]
	res = b
	for i := 0; i < len(a); i++ {
		res[i] -= a[i] * k
	}
	return res

}

// приведение  треугольному типу
func To_triangular(arr [][]float64) [][]float64 {
	if arr[0][0] == 0 {
		for i := 1; i < len(arr); i++ {
			if arr[i][0] != 0 {
				arr = Rep_r(0, i, arr)
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			arr[j] = SubRows(arr[i], arr[j])
		}
	}
	return arr
}

func Print_Arr(a [][]float64) {
	for _, i := range a {
		fmt.Println(i)
	}
}
