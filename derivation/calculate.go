package derivation

func FirstDeriv(infunc func(float64) float64, x float64) float64 {
	h := 0.1
	return (infunc(x+h) - infunc(x-h)) / (2 * h)
}

func SecondDeriv(infunc func(float64) float64, x float64) float64 {
	h := 0.1
	return (infunc(x+h) + infunc(x-h) - 2*infunc(x)) / (h * h)
}
