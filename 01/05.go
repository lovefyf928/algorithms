package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	file, err := os.Create("test.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	err1 := jpeg.Encode(file, m, &jpeg.Options{100})
	if err1 != nil {
		fmt.Println(err1)
	}
}
