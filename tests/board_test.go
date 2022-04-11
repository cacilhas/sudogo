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
	t.Run("Load", func(t *testing.T) {
		t.Run("right input", func(t *testing.T) {
			input := `+---+---+---+
|8..|...|..4|
|.5.|...|.82|
|...|.7.|.93|
+---+---+---+
|...|...|825|
|...|.9.|6..|
|46.|58.|..1|
+---+---+---+
|5..|.3.|4.8|
|1..|..7|.59|
|7.8|925|...|
+---+---+---+
`
			board, err := sudoku.LoadBoard(input)
			if err != nil {
				t.Fatal(err)
			}
			if got := board.String(); got != input {
				t.Fatalf("expected:\n%s\ngot:\n%s", input, got)
			}
		})
		t.Run("too small input", func(t *testing.T) {
			board, err := sudoku.LoadBoard("1234")
			if board != nil {
				t.Fatalf("expected nil, got:\n%s", board)
			}
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
		t.Run("too long input", func(t *testing.T) {
			input := ""
			for i := 0; i < 100; i++ {
				input += "0"
			}
			board, err := sudoku.LoadBoard(input)
			if board != nil {
				t.Fatalf("expected nil, got:\n%s", board)
			}
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	})
}
