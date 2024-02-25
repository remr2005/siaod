package slau

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
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

// решу через gonum
// https://ru.wikipedia.org/wiki/%D0%9C%D0%B5%D1%82%D0%BE%D0%B4_%D1%81%D0%BE%D0%BF%D1%80%D1%8F%D0%B6%D1%91%D0%BD%D0%BD%D1%8B%D1%85_%D0%B3%D1%80%D0%B0%D0%B4%D0%B8%D0%B5%D0%BD%D1%82%D0%BE%D0%B2_(%D0%B4%D0%BB%D1%8F_%D1%80%D0%B5%D1%88%D0%B5%D0%BD%D0%B8%D1%8F_%D0%A1%D0%9B%D0%90%D0%A3)
func MatScale(a float64, A *mat.Dense) *mat.Dense {
	r, c := A.Dims()
	buf := mat.NewDense(r, c, nil)
	buf.Scale(a, A)
	return buf
}

func MatProd(A *mat.Dense, B *mat.Dense) *mat.Dense {
	r, _ := A.Dims()
	_, c := B.Dims()
	buf := mat.NewDense(r, c, nil)
	buf.Product(A, B)
	return buf
}

func MatSub(A *mat.Dense, B *mat.Dense) *mat.Dense {
	r, c := A.Dims()
	buf := mat.NewDense(r, c, nil)
	buf.Sub(A, B)
	return buf
}

func MatAdd(A *mat.Dense, B *mat.Dense) *mat.Dense {
	r, c := A.Dims()
	buf := mat.NewDense(r, c, nil)
	buf.Add(A, B)
	return buf
}

func MatScalar(a []float64) float64 {
	res := 0.0
	for _, i := range a {
		res += i * i
	}
	return res
}

func InverseGradient(a [][]float64, b []float64, e float64) []float64 {
	A := mat.NewDense(len(a), len(a[0]), Two2one(a))
	B := mat.NewDense(len(b), 1, b)
	x := make([]float64, 0)
	for i := 0; i < len(a[0]); i++ {
		x = append(x, 1)
	}
	X := mat.NewDense(len(b), 1, x)
	R := mat.NewDense(len(b), 1, b)
	R_ := mat.DenseCopyOf(R)
	fmt.Println(1)

	Buf := mat.NewDense(len(a), 1, nil)
	fmt.Println(1)
	Buf.Product(A, X)
	fmt.Println(B.Dims())
	fmt.Println(Buf.Dims())
	R.Sub(B, Buf)

	alpha := 0.0
	var betta float64
	fmt.Println(1)
	Z := mat.DenseCopyOf(R)
	for MatScalar(R.RawMatrix().Data)/MatScalar(B.RawMatrix().Data) >= e {
		fmt.Println(1)

		Buf := mat.NewDense(1, 1, nil)
		Buf2 := mat.NewDense(1, 1, nil)
		Buf3 := mat.NewDense(len(a), 1, nil)
		Buf3.Product(A, Z)
		Buf2.Product(Buf3, Z)
		Buf.Product(R.T(), R)
		alpha = Buf.At(0, 0) / Buf2.At(0, 0)

		fmt.Println(1)

		Buf = mat.DenseCopyOf(Z)
		Buf.Scale(alpha, Z)
		X.Add(X, Buf)

		fmt.Println(1)

		Buf = mat.NewDense(len(a), 1, nil)
		Buf.Product(A, Z)
		Buf.Scale(alpha, Buf)
		R_ = mat.DenseCopyOf(R)
		R.Sub(R, Buf)

		fmt.Println(1)

		Buf = mat.NewDense(1, 1, nil)
		Buf2 = mat.NewDense(1, 1, nil)
		Buf.Product(R.T(), R)
		Buf2.Product(R_.T(), R_)
		betta = Buf.At(0, 0) / Buf2.At(0, 0)

		fmt.Println(1)

		buf := mat.DenseCopyOf(Z)
		buf.Scale(betta, buf)
		Z.Add(R, buf)

	}
	return X.RawMatrix().Data
}
