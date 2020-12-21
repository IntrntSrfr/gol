package main

import (
	"fmt"
	"image/gif"
	"os"
	"strings"
	"time"
)

func main() {
	grid := NewGrid(128)

	grid.Place(51, 51, pulsar)
	grid.Set(51, 51, 1)

	grid.Show()

	// this will make a gif, useful if theres a large grid
	render := &gif.GIF{}
	Iterate(&grid, 1000, false, true, render)
	SaveGif(render)

	// alternatively, this will display it in the console window
	// Iterate(&grid, 1000, true, false, nil, nil)
}
func SaveGif(g *gif.GIF) {
	f, _ := os.Create("./out.gif")
	gif.EncodeAll(f, g)
}

func Iterate(g *Grid, steps int, show, makeGif bool, gif *gif.GIF) {
	for i := 0; i < steps; i++ {
		if i%100 == 0 {
			fmt.Println(i)
		}
		*g = g.Step()
		if makeGif {
			gif.Image = append(gif.Image, newFrame(*g))
			gif.Delay = append(gif.Delay, 10)
		}
		if show {
			g.Show()
			time.Sleep(time.Millisecond * 150)
		}
	}
}

type Grid [][]int

func (g Grid) Place(dx, dy int, p Pattern) {
	if dx < 0 || dy < 0 || dy+len(p) > len(g) || dx+len(p[0]) > len(g[0]) {
		return
	}

	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			g.Set(dx+x, dy+y, p[y][x])
		}
	}
}

func (g Grid) Step() Grid {
	nGrid := NewGrid(len(g))

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			n := g.Neighbours(x, y)
			at := g.At(x, y)

			if at == 1 && n < 2 {
				continue
			}
			if at == 1 && n >= 2 && n <= 3 {
				nGrid.Set(x, y, 1)
			}
			if at == 1 && n > 3 {
				continue
			}
			if at == 0 && n == 3 {
				nGrid.Set(x, y, 1)
			}
		}
	}
	return nGrid
}

func NewGrid(size int) Grid {

	grid := make([][]int, size)
	for i := range grid {
		grid[i] = make([]int, size)
	}
	return grid
}

func (g Grid) Neighbours(x, y int) int {

	count := 0

	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			count += g.At(x+dx, y+dy)
		}
	}
	return count
}

func (g Grid) Set(x, y, state int) {
	if x < 0 || y < 0 || x >= len(g) || y >= len(g[x]) {
		return
	}

	if state != 0 && state != 1 {
		panic("state can only be 0 or 1")
	}
	g[y][x] = state
}
func (g Grid) At(x, y int) int {
	if x < 0 || y < 0 || x >= len(g) || y >= len(g[x]) {
		return 0
	}

	return g[y][x]
}
func (g Grid) Show() {

	sb := strings.Builder{}
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			sb.WriteString(map[int]string{0: ". ", 1: "# "}[g[y][x]])
		}
		sb.WriteRune('\n')
	}
	sb.WriteRune('\n')
	fmt.Println(sb.String())
}
