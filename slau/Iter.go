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


// решу через gonum
// https://ru.wikipedia.org/wiki/%D0%9C%D0%B5%D1%82%D0%BE%D0%B4_%D1%81%D0%BE%D0%BF%D1%80%D1%8F%D0%B6%D1%91%D0%BD%D0%BD%D1%8B%D1%85_%D0%B3%D1%80%D0%B0%D0%B4%D0%B8%D0%B5%D0%BD%D1%82%D0%BE%D0%B2_(%D0%B4%D0%BB%D1%8F_%D1%80%D0%B5%D1%88%D0%B5%D0%BD%D0%B8%D1%8F_%D0%A1%D0%9B%D0%90%D0%A3)
func InverseGradient(a [][]float64, b []float64, e float64) []float64 {
	A := mat.NewDense(len(a), len(a[0]), Two2one(a))
	R := mat.NewDense(len(b), 1, b)
	x := make([]float64, 0)
	for i := 0; i < len(a[0]); i++ {
		x = append(x, 1)
	}
	X := mat.NewDense(len(b), 1, x)
	B := mat.NewDense(len(a), 1, nil)
	B.Product(A, X)
	R.Sub(R, B)
	R_last := mat.DenseCopyOf(R)
	Z := mat.DenseCopyOf(R)
	var Alpha float64
	var Beta float64
	var calculate_alpha = func() float64 {
		Buf := mat.NewDense(1, 1, nil)
		Buf2 := mat.NewDense(1, 1, nil)
		Buf3 := mat.NewDense(len(a), 1, nil)
		Buf3.Product(A, Z)
		Buf.Product(R.T(), R)
		Buf2.Product(Buf3.T(), Z)
		return Buf.RawMatrix().Data[0] / Buf2.RawMatrix().Data[0]
	}
	var calculate_x = func() *mat.Dense {
		Buf := mat.DenseCopyOf(Z)
		Buf.Scale(Alpha, Buf)
		X.Add(X, Buf)
		return X
	}
	var calculate_r = func() *mat.Dense {
		Buf := mat.NewDense(len(a), 1, nil)
		Buf.Product(A, Z)
		Buf.Scale(Alpha, Buf)
		R_last = mat.DenseCopyOf(R)
		R.Sub(R, Buf)
		return R
	}
	var calculate_betta = func() float64 {
		Buf1 := mat.NewDense(1, 1, nil)
		Buf2 := mat.NewDense(1, 1, nil)
		Buf1.Product(R.T(), R)
		Buf2.Product(R_last.T(), R_last)
		return Buf1.RawMatrix().Data[0] / Buf2.RawMatrix().Data[0]
	}
	var calculate_z = func() *mat.Dense {
		Buf1 := mat.DenseCopyOf(Z)
		Buf1.Scale(Beta, Buf1)
		Z.Add(R, Buf1)
		return Z
	}
	/*
		var mod_vect = func() float64 {
			var (
				res1 float64
				res2 float64
			)
			for _, i := range R.RawMatrix().Data {
				res1 += i * i
			}
			for _, i := range b {
				res2 += i * i
			}
			return math.Sqrt(res1) / math.Sqrt(res2)
		}*/
	for i := 0.0; i < e; i++ {
		fmt.Println(Alpha, Beta)
		Alpha = calculate_alpha()
		X = calculate_x()
		R = calculate_r()
		Beta = calculate_betta()
		Z = calculate_z()
	}
	return X.RawMatrix().Data
}
