package main

import "golang.org/x/tour/pic"
import "image/color"
import "image"

type Image struct {
	w int
	h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	imgFunc := func(x, y int) uint8 {
		return uint8(x ^ y)
	}

	v := imgFunc(x, y)

	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{255, 70}
	pic.ShowImage(m)
}
