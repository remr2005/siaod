package integrals

import (
	"image/png"
	"math"
	"math/rand"
	"os"
)

func MonteKarlo(inFunc func(x float64) float64, a, b, n float64, iter int) (float64, float64) {
	var res float64
	var u float64 = rand.Float64()*(b-a) + a
	arr := make([]float64, 0)
	for j := 0; j < iter; j++ {
		for i := 1.0; i < n; i++ {
			res += inFunc(u)
			u = rand.Float64()*math.Abs(a-b) + a
		}
		arr = append(arr, (res*(b-a))/n)
		res = 0
	}
	return arr[0], Sred(arr)
}

func MonteKarloImage2() (int, int) {
	// чтение изображения
	existingImageFile, _ := os.Open("test4.png")
	// открытие изображения
	loadedImage, _ := png.Decode(existingImageFile)
	//массив куда будем записывать все итоги вычисления
	arr := make([]int, 0)
	// анонимная функция возвращающая площадь в метрах кв
	var calculate = func() int {
		var k int // кол-во попаданий
		// rgba код голубого цвета
		var c1 uint32 = 144
		var c2 uint32 = 218
		var c3 uint32 = 238
		var c4 uint32 = 255
		// кол-во бросков
		n := 10000000
		for i := 0; i < n; i++ {
			// генерацмя рандомных x,y
			x := rand.Intn(270)
			y := rand.Intn(449)
			// загрузка значений rgba пикселя с координатами (x,y)
			g1, g2, g3, g4 := loadedImage.At(x, y).RGBA()
			// проверка совпадания цвета
			if (c1 == g1%256) && (c2 == g2%256) && (c3 == g3%256) && (c4 == g4%256) {
				k += 1
			}
		}
		//вычисление площади в пикселях
		k = (k * 270 * 449) / n

		// в гугле с помощью линецки провел 6 км, вырезал прямоугольник с ними
		// ширина изображения 270 px. Длина одной стороны пикселя 6000/270=22.222222...
		// площадь одного пикселя в метрах 22.22222...*22.22222...=494 м кв
		return k * 494
	}
	// Запускаем алгоритм несколько раз
	for i := 0; i < 10; i++ {
		arr = append(arr, calculate())
	}
	// возвращаем площадь и погрешность
	return arr[0], SredInt(arr)
}

func MonteKarloImage3() (int, int) {
	existingImageFile, err := os.Open("test3.png")
	if err != nil {
		panic(err)
	}
	defer existingImageFile.Close()

	existingImageFile.Seek(0, 0)

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		panic(err)
	}
	// 387 h-475
	//270 449 test2
	// {144 218 238 255}
	arr := make([]int, 0)
	var calculate = func() int {
		var k int
		var c1 uint32 = 255
		var c2 uint32 = 0
		var c3 uint32 = 0
		var c4 uint32 = 255
		n := 10000000
		for i := 0; i < n; i++ {
			x := rand.Intn(840)
			y := rand.Intn(286)
			// fmt.Println()
			g1, g2, g3, g4 := loadedImage.At(x, y).RGBA()
			// fmt.Println(g1%256, g2%256, g3%256, g4%256, "-", loadedImage.At(x, y))
			if (c1 == g1%256) && (c2 == g2%256) && (c3 == g3%256) && (c4 == g4%256) {
				k += 1
			}
		}
		k = (k * 840 * 286) / n
		return k * 13
	}
	for i := 0; i < 10; i++ {
		arr = append(arr, calculate())
	}
	return arr[0], SredInt(arr)
}
