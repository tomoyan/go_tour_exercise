/*
Exercise: Images
Remember the picture generator you wrote earlier?
Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color;
the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
*/
package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w int
	h int
	v uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(w int, h int) color.Color {
	return color.RGBA{i.v, i.v, 200, 110}
}

func main() {
	m := Image{100, 100, 2}
	pic.ShowImage(m)
}
