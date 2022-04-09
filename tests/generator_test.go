package tests

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/cacilhas/sudogo/sudoku"
)

func TestGenerator(t *testing.T) {
	t.Run("HideCells", func(t *testing.T) {
		t.Run("Easy", func(t *testing.T) {
			rand.Seed(-1)
			board := sudoku.NewBoard()
			sudoku.Generator.HideCells(board, sudoku.EASY)
			expected := `+---+---+---+
|1..|.56|789|
|.5.|.8.|12.|
|789|123|456|
+---+---+---+
|.34|5..|..1|
|5.7|..1|2..|
|8.1|..4|.67|
+---+---+---+
|3..|.78|9.2|
|6.8|91.|34.|
|9.2|.4.|.78|
+---+---+---+
`
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%s\ngot:\n%s", expected, got)
			}
		})
		t.Run("Hard", func(t *testing.T) {
			rand.Seed(-1)
			board := sudoku.NewBoard()
			sudoku.Generator.HideCells(board, sudoku.HARD)
			expected := `+---+---+---+
|1..|.5.|789|
|.5.|.8.|12.|
|...|.23|.56|
+---+---+---+
|.3.|...|...|
|...|...|2..|
|..1|...|..7|
+---+---+---+
|3.5|...|9.2|
|6.8|91.|.4.|
|9.2|.4.|.78|
+---+---+---+
`
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%s\ngot:\n%s", expected, got)
			}
		})
	})

	t.Run("SwapDigits", func(t *testing.T) {
		rand.Seed(-1)
		board := sudoku.NewBoard()
		first := board.String()
		sudoku.Generator.SwapDigits(board)
		second := board.String()
		if first == second {
			t.Fatalf("expected different boards, got:\n%s", first)
		}
		checkBoard(board, t)
	})
	t.Run("SwapRows", func(t *testing.T) {
		rand.Seed(-1)
		board := sudoku.NewBoard()
		first := board.String()
		sudoku.Generator.SwapRows(board)
		second := board.String()
		if first == second {
			t.Fatalf("expected different boards, got:\n%s", first)
		}
		checkBoard(board, t)
	})
	t.Run("SwapColumns", func(t *testing.T) {
		rand.Seed(-1)
		board := sudoku.NewBoard()
		first := board.String()
		sudoku.Generator.SwapColumns(board)
		second := board.String()
		if first == second {
			t.Fatalf("expected different boards, got:\n%s", first)
		}
		checkBoard(board, t)
	})

	// ---------------------------------------------------------------------------
	t.Run("Generate", func(t *testing.T) {
		t.Run("Easy", func(t *testing.T) {
			rand.Seed(-1)
			board := sudoku.NewBoard()
			sudoku.Generator.Generate(board, sudoku.EASY)
			expected := `+---+---+---+
|8..|35.|164|
|95.|416|.82|
|...|278|.93|
+---+---+---+
|...|..4|825|
|.8.|.9.|6..|
|46.|58.|..1|
+---+---+---+
|5..|63.|4.8|
|1..|847|.59|
|748|925|.1.|
+---+---+---+
`
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%s\ngot:\n%s", expected, got)
			}
			expected = "9731[.]"
			if got := board.Get(1, 3).Debug(); got != expected {
				t.Fatalf("expected %s, got %s", expected, got)
			}
		})
		t.Run("Hard", func(t *testing.T) {
			rand.Seed(-1)
			board := sudoku.NewBoard()
			sudoku.Generator.Generate(board, sudoku.HARD)
			expected := `+---+---+---+
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
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%s\ngot:\n%s", expected, got)
			}
			expected = "93[.]"
			if got := board.Get(0, 3).Debug(); got != expected {
				t.Fatalf("expected %s, got %s", expected, got)
			}
		})
	})
}

func checkBoard(board sudoku.Board, t *testing.T) {
	checkRows(board, t)
	checkColumns(board, t)
	checkBlocks(board, t)
}

func checkColumns(board sudoku.Board, t *testing.T) {
	for x := 0; x < 9; x++ {
		t.Run(fmt.Sprintf("column %d", x), func(t *testing.T) {
			for y := 0; y < 9; y++ {
				cellValue := board.Get(x, y).Value()
				if cellValue == 0 {
					continue
				}
				for ly := 0; ly < 9; ly++ {
					if ly == y {
						continue
					}
					if cellValue == board.Get(x, ly).Value() {
						t.Fatalf("expected different values %d/%d, got:\n%s", y, ly, board)
					}
				}
			}
		})
	}
}

func checkRows(board sudoku.Board, t *testing.T) {
	for y := 0; y < 9; y++ {
		t.Run(fmt.Sprintf("row %d", y), func(t *testing.T) {
			for x := 0; x < 9; x++ {
				cellValue := board.Get(x, y).Value()
				if cellValue == 0 {
					continue
				}
				for lx := 0; lx < 9; lx++ {
					if lx == x {
						continue
					}
					if cellValue == board.Get(lx, y).Value() {
						t.Fatalf("expected different values %d/%d, got:\n%s", x, lx, board)
					}
				}
			}
		})
	}
}

func checkBlocks(board sudoku.Board, t *testing.T) {
	for b := 0; b < 9; b++ {
		bx := b % 3
		by := b / 3
		t.Run(fmt.Sprintf("block %d", b), func(t *testing.T) {
			for i := 0; i < 9; i++ {
				cellValue := board.Get(bx*3+i%3, by*3+i/3).Value()
				if cellValue == 0 {
					continue
				}
				for j := 0; j < 9; j++ {
					if j == i {
						continue
					}
					if cellValue == board.Get(bx*3+j%3, by*3+j/3).Value() {
						t.Fatalf("expected different values %d/%d, got:\n%s", i, j, board)
					}
				}
			}
		})
	}
}
