package interpolation

import (
	"math"
	"siaodMath/integrals"

	"gonum.org/v1/gonum/integrate"
)

// на равномерной сетке (P.S он не работает)
func FurieSet(f func(float64) float64, N float64, x []float64) []float64 {
	a0 := 0.0
	x_ := make([]float64, 0)
	for i := 1.0; i <= 2*N+1; i++ {
		res := (2 * math.Pi * (i - 1)) / (2*N + 1)
		x_ = append(x_, res)
		a0 += f(res)
	}
	a0 /= (2*N + 1)
	a := make([]float64, 0)
	b := make([]float64, 0)
	for k := 1.0; k < N+1; k++ {
		i_int := 0
		res_c := 0.0
		res_s := 0.0
		for i := 1.0; i <= 2*N+1; i++ {
			res_c += f(x_[i_int]) * math.Cos((2*k*i*math.Pi)/(2*N+1))
			res_s += f(x_[i_int]) * math.Sin((2*k*i*math.Pi)/(2*N+1))
			i_int += 1
		}
		a = append(a, res_c/(2*N+1))
		b = append(b, res_s/(2*N+1))
	}

	res := make([]float64, 0)
	for _, i := range x {
		sum := 0.0
		k := 1.0
		for j := 0; j < len(a); j++ {
			sum += a[j]*math.Cos(k*i) + b[j]*math.Sin(k*i)
			k += 1
		}
		res = append(res, a0/2+sum)
	}
	return res
}

func Furie(f func(float64) float64, a, b, N float64, x []float64) []float64 {
	a0, _, _ := integrals.Simpson(f, a, b, 10, 1000000000000, 0.001)
	a0 *= (1 / (b - a))
	a_ar := make([]float64, 0)
	b_ar := make([]float64, 0)
	for k := 1.0; k < N+1; k++ {
		var b_c = func(x float64) float64 {
			return f(x) * math.Cos(2*math.Pi*k*x/(b-a))
		}
		var b_s = func(x float64) float64 {
			return f(x) * math.Sin(2*math.Pi*k*x/(b-a))
		}
		res_c, _, _ := integrals.Simpson(b_c, a, b, 10, 100000000, 0.001)
		res_s, _, _ := integrals.Simpson(b_s, a, b, 10, 100000000, 0.001)
		a_ar = append(a_ar, res_c/(b-a))
		b_ar = append(b_ar, res_s/(b-a))
	}
	res := make([]float64, 0)
	for _, i := range x {
		sum := 0.0
		k := 1.0
		for j := 0; j < len(a_ar); j++ {
			sum += a_ar[j]*math.Cos(2*math.Pi*k*i/(b-a)) + b_ar[j]*math.Sin(2*math.Pi*k*i/(b-a))
			k += 1
		}
		res = append(res, a0/2+sum)
	}
	return res
}

func Furie1(f func(float64) float64, a, b, N float64, x []float64) []float64 {
	a0, _, _ := integrals.Simpson(f, a, b, 10, 1000000000000, 0.001)
	a0 *= (1 / (b - a))
	a_ar := make([]float64, 0)
	b_ar := make([]float64, 0)
	for k := 1.0; k < N+1; k++ {
		var b_c = func(x float64, k float64) float64 {
			return f(x) * math.Cos(2*math.Pi*k*x/(b-a))
		}
		var b_s = func(x float64, k float64) float64 {
			return f(x) * math.Sin(2*math.Pi*k*x/(b-a))
		}
		res_c, _, _ := integrals.Simpson(func(x float64) float64 { return b_c(x, k) }, a, b, 10, 100000000, 0.001)
		res_s, _, _ := integrals.Simpson(func(x float64) float64 { return b_s(x, k) }, a, b, 10, 100000000, 0.001)
		a_ar = append(a_ar, res_c/(b-a))
		b_ar = append(b_ar, res_s/(b-a))
	}
	res := make([]float64, 0)
	for _, i := range x {
		sum := 0.0
		for j := 0; j < len(a_ar); j++ {
			sum += a_ar[j]*math.Cos(2*math.Pi*float64(j+1)*i/(b-a)) + b_ar[j]*math.Sin(2*math.Pi*float64(j+1)*i/(b-a))
		}
		res = append(res, a0/2+sum)
	}
	return res
}

func Furie2(f func(float64) float64, a, b, N float64, x []float64) []float64 {
	cof := make([]float64, 0)
	a0, _, _ := integrals.Simpson(f, a, b, 10, 1000000000000, 0.001)
	a0 *= (1 / (b))
	cof = append(cof, a0)
	for k := 1.0; k < N+1; k++ {
		var b_c = func(x float64) float64 {
			return f(x) * math.Cos(2*math.Pi*k*x/(b-a))
		}
		var b_s = func(x float64) float64 {
			return f(x) * math.Sin(2*math.Pi*k*x/(b-a))
		}
		x_ := make([]float64, 0)
		f_c := make([]float64, 0)
		f_s := make([]float64, 0)
		for g := a; g < b; g += 0.01 {
			x_ = append(x_, g)
			f_c = append(f_c, b_c(g))
			f_s = append(f_s, b_s(g))
		}
		x_ = append(x_, b)
		f_c = append(f_c, b_c(b))
		f_s = append(f_s, b_s(b))
		res_c := integrate.Simpsons(x_, f_c)
		res_s := integrate.Simpsons(x_, f_s)
		cof = append(cof, res_c/(b))
		cof = append(cof, res_s/(b))
	}
	res := make([]float64, 0)
	for _, i := range x {
		sum := 0.0
		for j := 1; j < len(cof)-1; j++ {
			sum += cof[j]*math.Cos(2*math.Pi*float64(j)*i/(b-a)) + cof[j+1]*math.Sin(2*math.Pi*float64(j)*i/(b-a))
		}
		res = append(res, cof[0]/2+sum)
	}
	return res
}
