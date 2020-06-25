// Create an implementation of image.Image

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

//
type Image struct {
	width, height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	z := uint8(x ^ y)
	return color.RGBA{z, z, 255, 255}
}

func main() {
	m := Image{500, 500}
	pic.ShowImage(m)
}