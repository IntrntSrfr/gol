package main

import (
	"fmt"
	"image/gif"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {

	cpuF, _ := os.Create("cpu.prof")
	memF, _ := os.Create("mem.prof")

	pprof.StartCPUProfile(cpuF)
	defer pprof.StopCPUProfile()

	seed := time.Now().Unix()

	grid := NewGrid(32, 32, seed, true)
	bufGrid := NewGrid(32, 32, seed, true)

	rand.Seed(seed)

	l := (grid.h * grid.w) / 1
	for l > 0 {
		grid.Set(rand.Intn(grid.w), rand.Intn(grid.h), 1)
		l--
	}

	grid.DeepCopy(bufGrid)
	//render := &gif.GIF{}

	for i := 0; i < 1000; i++ {
		grid.Step(bufGrid)
		grid.DeepCopy(bufGrid)

		//render.Image = append(render.Image, newFrame(grid))
		//render.Delay = append(render.Delay, 10)
		grid.Show()
		time.Sleep(time.Millisecond * 150)
	}

	grid.Show()
	//SaveGif(render)

	// this will make a gif, useful if theres a large grid
	/*
			render := &gif.GIF{}
		Iterate(grid, 1000, false, true, render)
		err := SaveGif(render)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	// alternatively, this will display it in the console window
	//grid.Show()
	//Iterate(grid, 100, false, false, nil)

	runtime.GC()
	pprof.WriteHeapProfile(memF)

}
func SaveGif(g *gif.GIF) error {
	f, err := os.Create("./out.gif")
	if err != nil {
		return err
	}
	return gif.EncodeAll(f, g)
}

/*
func Iterate(g *Grid, steps int, show, makeGif bool, gif *gif.GIF) {
	for i := 0; i < steps; i++ {
		g = g.Step()
		if makeGif {
			if i%(steps/100) == 0 {
				fmt.Println(i)
			}
			gif.Image = append(gif.Image, newFrame(g))
			gif.Delay = append(gif.Delay, 10)
		}
		if show {
			g.Show()
			time.Sleep(time.Millisecond * 150)
		}
	}
}*/

type Grid struct {
	data [][]int
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
			/*
				if at == 1 && n < 2 {
					g.Set(x, y, 0)
				}
			*/
			if at == 1 && (n != 2 && n != 3) {
				g.Set(x, y, 0)
			}
			/*
				if at == 1 && n > 3 {
					g.Set(x, y, 0)
				}
			*/
			if at == 0 && n == 3 {
				g.Set(x, y, 1)
			}
		}
	}
	g.gen = g.gen + 1
}

func NewGrid(h, w int, seed int64, wrap bool) *Grid {

	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
	}

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
			dst.data[y][x] = g.data[y][x]
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

	g.data[y][x] = state
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

	return g.data[y][x]
}

func (g *Grid) Show() {
	/*
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	*/

	index := 0
	buf := make([]byte, g.h*g.w*2+g.h+256)

	inf := fmt.Sprintf("seed:       %v\ngeneration: %v\n", g.seed, g.gen)
	for range inf {
		buf[index] = inf[index]
		index++
	}

	//fmt.Println(index)

	//w := bufio.NewWriterSize(os.Stdout, len(buf))

	//w.WriteString(fmt.Sprintf("seed:       %v\ngeneration: %v\n", g.seed, g.gen))
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			s := g.data[y][x]
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

	os.Stdout.Write(buf[:index])
	//os.Stdout.WriteString(fmt.Sprint(buf))

	//w.Flush()

	//fmt.Println(sb.Len())

	//fmt.Println(sb)
	//os.Stdout.WriteString(sb.String())
	//fmt.Println(sb.String())
}
