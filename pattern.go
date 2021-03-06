package gol

type Pattern [][]int

var (

	// # #
	// # #
	block Pattern = [][]int{
		{1, 1},
		{1, 1},
	}

	// . # # .
	// # . . #
	// . # # .
	beehive Pattern = [][]int{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
	}

	// . # # .
	// # . . #
	// . # . #
	// . . # .
	loaf Pattern = [][]int{
		{0, 1, 1, 0},
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 0},
	}

	// # # .
	// # . #
	// . # .
	boat Pattern = [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
	}

	// . # .
	// # . #
	// . # .
	tub Pattern = [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
	}

	// # # .
	// # . #
	// . # #
	ship Pattern = [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, 1},
	}

	// # # #
	blinker Pattern = [][]int{
		{1, 1, 1},
	}

	// . # # #
	// # # # .
	toad Pattern = [][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 0},
	}

	// # # . .
	// # # . .
	// . . # #
	// . . # #
	beacon Pattern = [][]int{
		{1, 1, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 1, 1},
		{0, 0, 1, 1},
	}

	// . . # # # . . . # # # . .
	// . . . . . . . . . . . . .
	// # . . . . # . # . . . . #
	// # . . . . # . # . . . . #
	// # . . . . # . # . . . . #
	// . . # # # . . . # # # . .
	// . . . . . . . . . . . . .
	// . . # # # . . . # # # . .
	// # . . . . # . # . . . . #
	// # . . . . # . # . . . . #
	// # . . . . # . # . . . . #
	// . . . . . . . . . . . . .
	// . . # # # . . . # # # . .
	pulsar Pattern = [][]int{
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
	}

	// . . #
	// # . #
	// . # #
	glider Pattern = [][]int{
		{0, 0, 1},
		{1, 0, 1},
		{0, 1, 1},
	}
)
