package slau

import "gonum.org/v1/gonum/mat"

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

func GausZeidel(a [][]float64, b []float64) {
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
