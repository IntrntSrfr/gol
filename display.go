package main

import (
	"image"
	"image/color"
)

var palette = []color.Color{
	color.Black,
	color.White,
}

func newFrame(g *Grid) *image.Paletted {
	frame := image.NewPaletted(image.Rect(0, 0, g.w, g.h), palette)

	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			if g.At(x, y) == 0 {
				frame.SetColorIndex(x, y, 0)
			} else {
				frame.SetColorIndex(x, y, 1)
			}
		}
	}
	return frame
}
