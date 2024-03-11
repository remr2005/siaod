package makefunctions

func BasicFunc(x float64) float64 {
	return x
}

func Addconst(f func(float64) float64, c float64) func(float64) float64 {
	return func(a float64) float64 { return f(a) + c }
}

func Prodconst(f func(float64) float64, c float64) func(float64) float64 {
	return func(a float64) float64 { return f(a) * c }
}

func AddFunc(f1 func(float64) float64, f2 func(float64) float64) func(float64) float64 {
	return func(a float64) float64 { return f1(a) + f2(a) }
}

func SubFunc(f1 func(float64) float64, f2 func(float64) float64) func(float64) float64 {
	return func(a float64) float64 { return f1(a) - f2(a) }
}

func ProdFunc(f1 func(float64) float64, f2 func(float64) float64) func(float64) float64 {
	return func(a float64) float64 { return f1(a) * f2(a) }
}

func DivFunc(f1 func(float64) float64, f2 func(float64) float64) func(float64) float64 {
	return func(a float64) float64 { return f1(a) / f2(a) }
}
