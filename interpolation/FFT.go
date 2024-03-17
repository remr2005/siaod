package interpolation

import (
	"math"
	"math/cmplx"
)

func FFT(x []complex128) []complex128 {
	N := len(x)
	N_com := complex(float64(N), 0)
	if N <= 1 {
		return x
	}
	even := make([]complex128, 0)
	odd := make([]complex128, 0)
	for i := 0; i < len(x); i++ {
		if i%2 == 0 {
			even = append(even, x[i])
		} else {
			odd = append(odd, x[i])
		}
	}
	result := make([]complex128, 0)
	even_fft := FFT(even)
	odd_fft := FFT(odd)
	factor := make([]complex128, 0)
	for k := 0; k < N/2; k++ {
		k_com := complex(float64(k), 0)
		exponent := cmplx.Exp(complex(0, -2) * complex(math.Pi, 0) * k_com / N_com)
		factor = append(factor, exponent*odd_fft[k])
		result = append(result, even_fft[k]+factor[k])
	}
	for k := 0; k < N/2; k++ {
		result = append(result, even_fft[k]-factor[k])
	}
	return result
}
