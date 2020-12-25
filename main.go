package main

import (
	"flag"
	"fmt"
	"image/gif"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	defer fmt.Print("\u001B[0m")
	display := flag.Bool("show", false, "display iterations")

	var iters int
	flag.IntVar(&iters, "iters", 1000, "how many iterations")

	var delay int
	flag.IntVar(&delay, "delay", 150, "delay per frame, useless if -show is not used")
	flag.Parse()

	out, _ := os.Create("./example.gif")

	NewGame(time.Now().Unix(), 85, 256, iters, delay, true, *display, out, 4)

}

func NewGame(seed int64, height, width, iters, delay int, wrap, display bool, out io.Writer, scale int) {
	grid := NewGrid(height, width, seed, wrap)
	bufGrid := NewGrid(height, width, seed, wrap)

	rand.Seed(seed)

	l := (grid.h * grid.w) / 4
	for l > 0 {
		grid.Set(rand.Intn(grid.w), rand.Intn(grid.h), 1)
		l--
	}
	grid.DeepCopy(bufGrid)

	var render *gif.GIF
	if out != nil {
		render = &gif.GIF{}
	}

	// scale of resulting gif

	if display {
		fmt.Print("\u001b[2J")
	}
	for i := 0; i < iters; i++ {

		grid.Step(bufGrid)
		grid.DeepCopy(bufGrid)

		if render != nil {
			render.Image = append(render.Image, newFrame(grid, scale))
			render.Delay = append(render.Delay, 10)
		}

		if display {
			grid.Show()
			time.Sleep(time.Millisecond * time.Duration(delay))
		}
	}
	if display {
		grid.Show()
	}

	SaveGif(render, out)

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

func NewGrid(h, w int, seed int64, wrap bool) *Grid {
	grid := make([]int, h*w)
	return &Grid{
		data: grid,
		h:    h,
		w:    w,
		seed: seed,
		wrap: wrap,
	}
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
