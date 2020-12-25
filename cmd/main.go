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

	var iters int
	flag.IntVar(&iters, "iters", 1000, "how many iterations")

	var delay int
	flag.IntVar(&delay, "delay", 150, "delay per frame, useless if -show is not used")
	flag.Parse()

	out, _ := os.Create("./out.gif")

	gol.NewGame(time.Now().Unix(), 85, 256, iters, delay, true, *display, out, 4)

}
