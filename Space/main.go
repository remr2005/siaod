package main

import (
	"fmt"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	existingImageFile, err := os.Open("test2.png")
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
	var k int
	var c1 uint32 = 144
	var c2 uint32 = 218
	var c3 uint32 = 238
	var c4 uint32 = 255
	n := 10000000
	for i := 0; i < n; i++ {
		x := rand.Intn(270)
		y := rand.Intn(449)
		// fmt.Println()
		g1, g2, g3, g4 := loadedImage.At(x, y).RGBA()
		// fmt.Println(g1%256, g2%256, g3%256, g4%256, "-", loadedImage.At(x, y))
		if (c1 == g1%256) && (c2 == g2%256) && (c3 == g3%256) && (c4 == g4%256) {
			k += 1
		}
	}
	k = (k * 270 * 449) / n

	fmt.Println(k * 494)
}

func second_attm() {
	existingImageFile, err := os.Open("test2.png")
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
	var k int
	var c1 uint32 = 144
	var c2 uint32 = 218
	var c3 uint32 = 238
	var c4 uint32 = 255
	n := 10000000
	for i := 0; i < n; i++ {
		x := rand.Intn(270)
		y := rand.Intn(449)
		// fmt.Println()
		g1, g2, g3, g4 := loadedImage.At(x, y).RGBA()
		// fmt.Println(g1%256, g2%256, g3%256, g4%256, "-", loadedImage.At(x, y))
		if (c1 == g1%256) && (c2 == g2%256) && (c3 == g3%256) && (c4 == g4%256) {
			k += 1
		}
	}
	k = (k * 270 * 449) / n

	fmt.Println(k * 494)
}
