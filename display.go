package gol

import (
	"image"
	"image/color"
	"image/gif"
	"io"
)

var palette = []color.Color{
	color.Black,
	color.White,
}

func newFrame(g *Grid, scale int) *image.Paletted {
	if scale <= 0 {
		panic("scale must be positive")
	}
	frame := image.NewPaletted(image.Rect(0, 0, g.w*scale, g.h*scale), palette)

	for y := 0; y < g.h*scale; y++ {
		for x := 0; x < g.w*scale; x++ {
			if g.At(x/scale, y/scale) == 0 {
				frame.SetColorIndex(x, y, 0)
			} else {
				frame.SetColorIndex(x, y, 1)
			}
		}
	}
	return frame
}

func SaveGif(g *gif.GIF, f io.Writer) error {
	return gif.EncodeAll(f, g)
}
