package gol

import "testing"

var (
	grid, _    = NewGrid(64, 64, 0, true)
	bufGrid, _ = NewGrid(64, 64, 0, true)
)

func BenchmarkGrid_Step(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid.Step(bufGrid)
		grid.DeepCopy(bufGrid)
	}
}

func BenchmarkGrid_DeepCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid.DeepCopy(bufGrid)
	}
}

func BenchmarkGrid_At(b *testing.B) {
	for i := 0; i < b.N; i++ {
		grid.At(32, 32)
	}
}
