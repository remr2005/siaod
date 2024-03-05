package derivation

func FirstDeriv(infunc func(float64) float64, x float64) float64 {
	h := 0.1
	return (infunc(x+h) - infunc(x-h)) / (2 * h)
}

func SecondDeriv(infunc func(float64) float64, x float64) float64 {
	h := 0.1
	return (infunc(x+h) + infunc(x-h) - 2*infunc(x)) / (h * h)
}

func FirstDerivArray(infunc func([]float64) float64, x []float64, n int) float64 {
	h := 0.1
	xCopy := make([]float64, len(x))
	copy(xCopy, x)
	xCopyCopy := make([]float64, len(xCopy))
	copy(xCopyCopy, xCopy)
	xCopyCopy[n] -= h
	xCopy[n] += h
	return (infunc(xCopy) - infunc(xCopyCopy)) / (2 * h)
}
