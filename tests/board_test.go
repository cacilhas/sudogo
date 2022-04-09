package tests

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/cacilhas/sudogo/sudoku"
)

func TestBoard(t *testing.T) {
	t.Run("NewBoard", func(t *testing.T) {
		board := sudoku.NewBoard()
		t.Run("String", func(t *testing.T) {
			expected := `+---+---+---+
|123|456|789|
|456|789|123|
|789|123|456|
+---+---+---+
|234|567|891|
|567|891|234|
|891|234|567|
+---+---+---+
|345|678|912|
|678|912|345|
|912|345|678|
+---+---+---+
`
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%s\ngot:\n%s", expected, got)
			}
		})
		t.Run("Get", func(t *testing.T) {
			for y := 0; y < 9; y++ {
				for x := 0; x < 9; x++ {
					t.Run(fmt.Sprintf("%d,%d", x, y), func(t *testing.T) {
						cell := board.Get(x, y)
						expected := fmt.Sprintf("%d[%d]", cell.Value(), cell.Value())
						if got := cell.Debug(); got != expected {
							t.Fatalf("expected %s, got %s", expected, got)
						}
					})
				}
			}
		})
	})
	t.Run("mutable", func(t *testing.T) {
		board := sudoku.NewBoard()
		for i := 0; i < 9; i++ {
			x := rand.Intn(9)
			y := rand.Intn(9)
			board.Get(x, y).Set(0)
			if board.Get(x, y).IsSet() {
				t.Fatalf("expected empty, got %s\n%s", board.Get(x, y), board)
			}
		}
	})
}
