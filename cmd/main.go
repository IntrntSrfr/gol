package main

import (
	"flag"
	"github.com/intrntsrfr/gol"
	"os"
	"time"
)

func main() {
	var iters, delay, height, width int
	var export string
	var seed int64

	show := flag.Bool("show", false, "display each iteration")
	skip := flag.Bool("skip", false, "show only the last iteration")

	flag.Int64Var(&seed, "seed", time.Now().Unix(), "generation seed")
	flag.StringVar(&export, "export", "", "export to gif")
	flag.IntVar(&iters, "iters", 1000, "how many iterations")
	flag.IntVar(&delay, "delay", 150, "delay per frame, useless if -show is not used")
	flag.IntVar(&height, "height", 32, "map height - min 16")
	flag.IntVar(&width, "width", 32, "map width - min 16")

	flag.Parse()

	// if its not going to export nor show, why bother?
	// also if its small its kinda shit
	if height < 16 || width < 16 || (!*show && export == "") {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var out *os.File
	if export != "" {
		out, _ = os.Create("./" + export)
	}

	game, _ := gol.NewGame(seed, height, width, true)
	game.Run(iters, delay, *show, *skip, export, 4)
	game.Export(out)
}
