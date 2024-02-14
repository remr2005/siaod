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

func ProdRow(a []float64, b []float64) []float64 {
	res := make([]float64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] * b[i]
	}
	return res
}

func Sum(a []float64) (res float64) {
	for _, i := range a {
		res += i
	}
	return res
}

// слайс в матрицу гонам
func Array_to_matrix(arr [][]float64) mat.Dense {
	return *mat.NewDense(len(arr), len(arr[0]), Two2one(arr))
}

// минусовать строки слайса
func SubRows(a, b []float64) (res []float64) {
	var nul int
	for n, i := range a {
		if i != 0 {
			nul = n
			break
		}
	}
	k := b[nul] / a[nul]
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

// вывести двумерный массив
func Print_Arr(a [][]float64) {
	for _, i := range a {
		fmt.Println(i)
	}
}

// решение по гаусу
func Gaus(a [][]float64, b []float64) ([]float64, error) {
	x := make([]float64, len(a[0]))
	for i := 0; i < len(x); i++ {
		x[i] = 1
	}
	arr, err := Conect(a, b)
	if err != nil {
		return []float64{}, err
	}
	arr = To_triangular(arr)
	arr_b := make([][]float64, 0)
	for _, i := range arr {
		var n int
		for j := 0; j < len(i)-1; j++ {
			if i[j] != 0 {
				n = j
				break
			}
		}
		arr_b = append(arr_b, Scale(1/i[n], i))
	}
	for i := len(a[0]) - 2; i >= 0; i-- {
		x[i] = FindX(arr_b[i], x)
	}
	return x, nil
}

// объединение матриц коофициентов и матрицы свободных членов
func Conect(a [][]float64, b []float64) ([][]float64, error) {
	if len(a) != len(b) {
		return [][]float64{}, fmt.Errorf("размер матрицы коофициентов не соответствует размеру матрицы свободных членов")
	}
	res := a
	for n, i := range b {
		res[n] = append(res[n], i)
	}
	return res, nil
}

func Scale(a float64, b []float64) (res []float64) {
	res = b
	for n, i := range b {
		res[n] = a * i
	}
	return res
}

func FindX(a []float64, b []float64) float64 {
	var eins int = -1
	for n, i := range a {
		if i == 1 {
			eins = n
			break
		}
	}
	x := a[len(a)-1]
	for i := eins + 1; i < len(a)-1; i++ {
		x -= a[i] * b[i]
	}
	return x
}
