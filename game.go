package gol

import (
	"errors"
	"fmt"
	"image/gif"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Game struct {
	grid    *Grid
	bufGrid *Grid
	render  *gif.GIF
	rng     *rand.Rand
}

func NewGame(seed int64, height, width int, wrap bool) (*Game, error) {
	rng := rand.New(rand.NewSource(seed))
	grid, err := NewGrid(height, width, seed, wrap)
	if err != nil {
		return nil, err
	}
	bufGrid, err := NewGrid(height, width, seed, wrap)
	if err != nil {
		return nil, err
	}

	for i := 0; i < (grid.h*grid.w)/6; i++ {
		grid.Set(rng.Intn(grid.w), rng.Intn(grid.h), 1)
	}
	grid.DeepCopy(bufGrid)

	return &Game{
		grid:    grid,
		bufGrid: bufGrid,
		rng:     rng,
	}, nil
}

func (g *Game) Export(out io.Writer) error {
	if g.render == nil {
		return errors.New("there is nothing to export")
	}
	return SaveGif(g.render, out)
}

func (g *Game) Run(iterations, delay int, show, skip bool, export string, scale int) {
	if export != "" {
		g.render = &gif.GIF{}
	}
	if show && !skip {
		fmt.Print("\u001b[2J")
	}
	for i := 0; i < iterations; i++ {
		g.grid.Step(g.bufGrid)
		g.grid.DeepCopy(g.bufGrid)
		if export != "" {
			g.render.Image = append(g.render.Image, newFrame(g.grid, scale))
			g.render.Delay = append(g.render.Delay, 10)
		}
		if show && !skip {
			g.grid.Show()
			time.Sleep(time.Millisecond * time.Duration(delay))
		}
	}
	if show {
		g.grid.Show()
		fmt.Print("\u001b[0m")
	}
}

type Grid struct {
	data []int
	h    int
	w    int
	gen  int
	seed int64
	wrap bool
}

func (g *Grid) Place(dx, dy int, p Pattern) {
	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			g.Set(dx+x, dy+y, p[y][x])
		}
	}
}

func (g *Grid) Step(src *Grid) {
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			n := src.Neighbours(x, y)
			at := src.At(x, y)
			if at == 1 && (n != 2 && n != 3) {
				g.Set(x, y, 0)
			}
			if at == 0 && n == 3 {
				g.Set(x, y, 1)
			}
		}
	}
	g.gen = g.gen + 1
}

func NewGrid(h, w int, seed int64, wrap bool) (*Grid, error) {
	if h <= 0 || w <= 0 {
		return nil, errors.New("Grid dimensions must be positive")
	}
	grid := make([]int, h*w)
	return &Grid{
		data: grid,
		h:    h,
		w:    w,
		seed: seed,
		wrap: wrap,
	}, nil
}

func (g *Grid) DeepCopy(dst *Grid) {
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			dst.data[y*g.w+x] = g.data[y*g.w+x]
		}
	}
}

func (g *Grid) Neighbours(x, y int) int {
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

func (g *Grid) Set(x, y, state int) {

	if g.wrap {

		if x < 0 {
			x = g.w + x
		}
		if y < 0 {
			y = g.h + y
		}
		if x >= g.w {
			x = x - g.w
		}
		if y >= g.h {
			y = y - g.h
		}

	} else {
		if x < 0 || y < 0 || x >= g.w || y >= g.h {
			return
		}
	}

	g.data[y*g.w+x] = state
}
func (g *Grid) At(x, y int) int {
	if g.wrap {

		if x < 0 {
			x = g.w + x
		}
		if y < 0 {
			y = g.h + y
		}
		if x >= g.w {
			x = x - g.w
		}
		if y >= g.h {
			y = y - g.h
		}
	} else {
		if x < 0 || y < 0 || x >= g.w || y >= g.h {
			return 0
		}
	}

	return g.data[y*g.w+x]
}

func (g *Grid) Show() {

	/*
		index := 0
		buf := make([]byte, g.h*g.w*2+g.h+256)

		inf := fmt.Sprintf("seed:       %v\ngeneration: %v\n", g.seed, g.gen)
		for range inf {
			buf[index] = inf[index]
			index++
		}
	*/
	//fmt.Println(index)

	//w := bufio.NewWriterSize(os.Stdout, len(buf))
	sb := strings.Builder{}

	fmt.Print("\u001b[2J")

	fmt.Print("\u001b[1;1H")
	//fmt.Print("\u001b[2K")
	fmt.Println(fmt.Sprintf("\u001B[0mseed:       %v\ngeneration: %v\n", g.seed, g.gen))

	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			s := g.data[y*g.w+x] //[x]
			if s == 0 {
				sb.WriteString("\u001b[34m. ")
			} else {
				sb.WriteString("\u001b[31m# ")
			}
		}
		sb.WriteRune('\n')
	}

	os.Stdout.WriteString(sb.String())

	/*
		for y := 0; y < g.h; y++ {
			for x := 0; x < g.w; x++ {
				s := g.data[y*g.w+x] //[x]
				if s == 0 {
					buf[index] = '.'
				} else {
					buf[index] = '#'
				}
				buf[index+1] = ' '
				index += 2
				//w.WriteString(map[int]string{0: ". ", 1: "# ", 2:"X "}[g.data[y][x]])
			}
			buf[index] = '\n'
			index++
			//w.WriteString("\n")
		}
		//fmt.Println(index)

		buf[index] = '\n'
		index++

		os.Stdout.Write(buf[:index])

	*/
	//os.Stdout.WriteString(fmt.Sprint(buf))

	//w.Flush()

	//fmt.Println(sb.Len())

	//fmt.Println(sb)
	//os.Stdout.WriteString(sb.String())
	//fmt.Println(sb.String())
}
