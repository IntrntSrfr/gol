package main

import (
	"flag"
	"fmt"
	"github.com/intrntsrfr/gol"
	"os"
	"time"
)

func main() {

	defer fmt.Print("\u001B[0m")
	display := flag.Bool("show", false, "display iterations")
	export := flag.Bool("export", false, "export to gif")

	var iters int
	flag.IntVar(&iters, "iters", 1000, "how many iterations")

	var delay int
	flag.IntVar(&delay, "delay", 150, "delay per frame, useless if -show is not used")

	var height, width int
	flag.IntVar(&height, "height", 32, "map height - min 16")
	flag.IntVar(&width, "width", 32, "map width - min 16")

	flag.Parse()

	if height < 16 || width < 16 || (!*display && !*export) {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var out *os.File
	if *export {
		out, _ = os.Create("./out.gif")
	}

	gol.NewGame(time.Now().Unix(), height, width, iters, delay, true, *display, out, 4)

}
