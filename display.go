package main

import (
	"image"
	"image/color"
)

var palette = []color.Color{
	color.Black,
	color.White,
}

func newFrame(grid Grid) *image.Paletted {
	frame := image.NewPaletted(image.Rect(0, 0, len(grid), len(grid[0])), palette)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid.At(x, y) == 0 {
				frame.SetColorIndex(x, y, 0)
			} else {
				frame.SetColorIndex(x, y, 1)
			}
		}
	}
	return frame
}
