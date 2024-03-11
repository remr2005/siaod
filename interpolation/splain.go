package interpolation

import (
	"fmt"
	"siaodMath/slau"

	"gonum.org/v1/gonum/mat"
)

// http://www.machinelearning.ru/wiki/index.php?title=%D0%98%D0%BD%D1%82%D0%B5%D1%80%D0%BF%D0%BE%D0%BB%D1%8F%D1%86%D0%B8%D1%8F_%D0%BA%D1%83%D0%B1%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B8%D0%BC%D0%B8_%D1%81%D0%BF%D0%BB%D0%B0%D0%B9%D0%BD%D0%B0%D0%BC%D0%B8
func Splain(f func(float64) float64, x []float64) {
	n := len(x)
	Mat := make([][]float64, n)
	B := make([]float64, n)
	h := make([]float64, n-1)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		Mat[i] = make([]float64, n)
	}
	for i := 0; i < n-1; i++ {
		h[i] = x[i+1] - x[i]
		y[i] = f(x[i])
		if i == n-2 {
			y[i+1] = f(x[i+1])
		}
	}
	fmt.Println(len(h))
	n = len(h)
	for i := 0; i < n-1; i++ {
		B[i] = 3 * ((y[i+2]-y[i+1])/h[i+1] - (y[i+1]-y[i])/h[i])
		Mat[i][i+1] = h[i+1]
		Mat[i][i] = 2 * (h[i] + h[i+1])
		if i > 0 {
			Mat[i][i-1] = h[i]
		}
		if i == n-2 {
			Mat[i+1][i+1] = 2 * h[i]
		}
	}
	n = len(x)
	A := mat.NewDense(n, n, slau.Two2one(Mat))
	A.Inverse(A)
	b := mat.NewDense(n, 1, B)
	b.Product(A, b)
	fmt.Println(1)
	arr_c := b.RawMatrix().Data
	arr_b := make([]float64, len(arr_c))
	arr_d := make([]float64, len(arr_c))
	fmt.Println(len(y), len(arr_c))
	for i := 1; i < len(arr_c)-1; i++ {
		arr_b[i] = (y[i]-y[i-1])/h[i] - (h[i]*(arr_c[i+1]+2*arr_c[i]))/3
		arr_d[i] = (arr_c[i+1] - arr_c[i]) / (3 * h[i])
		if i == len(arr_c)-2 {
			i += 1
			arr_b[i] = (y[i]-y[i-1])/h[i] - (2/3)*h[i]*arr_c[i]
			arr_d[i] = -(arr_c[i] / (3 * h[i]))
		}
	}
	fmt.Println("a:", y)
	fmt.Println("b:", arr_b)
	fmt.Println("c:", arr_c)
	fmt.Println("d:", arr_d)
}

func Splain(f func(float64) float64, x []float64) []float64 {
	n := len(x) - 1
	a := make([]float64, 0)
	b := make([]float64, 0)
	c := make([]float64, 0)
	d := make([]float64, 0)
	m := make([]float64, n+1)
	fmt.Println(n)
	h := make([]float64, n+1)
	y := make([]float64, n+1)
	y[n] = f(x[n])
	fmt.Println(1)
	for i := 0; i < n; i++ {
		y[i] = f(x[i])
		h[i] = x[i+1] - x[i]
	}
	fmt.Println(1, len(y), len(x))
	for i := 1; i <= n-1; i++ {
		a = append(a, (h[i]+h[i+1])/3)
		d = append(d, (y[i+1]-y[i])/h[i+1]-(y[i]-y[i-1])/h[i])
		b = append(b, h[i+1]/6)
		c = append(c, h[i]/6)
	}
	fmt.Println(1, len(c), len(b), len(a), len(d))
	l := make([]float64, 1)
	mu := make([]float64, 1)
	for i := 1; i <= n; i++ {
		m[i+1] = (d[i]*(-c[i])*mu[i-1] - (a[i]+c[i]*l[i-1])*m[i]) / b[i]
		mu = append(mu, (d[i]-c[i]*mu[i-1])/(a[i]+c[i]*l[i-1]))
		l = append(l, (-b[i])/(a[i]+c[i]*l[i-1]))
	}
	fmt.Println(1)
	return m
}

