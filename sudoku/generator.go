package sudoku

import (
	"math/rand"
)

type GeneratorType struct {
	Generate    func(Board, Level)
	HideCells   func(Board, Level)
	SwapDigits  func(Board)
	SwapRows    func(Board)
	SwapColumns func(Board)
}

var Generator GeneratorType

func init() {
	Generator.Generate = func(board Board, level Level) {
		Generator.SwapDigits(board)
		Generator.SwapRows(board)
		Generator.SwapColumns(board)
		Generator.HideCells(board, level)
		board.Fix()
	}

	Generator.HideCells = func(board Board, level Level) {
		level.Exec(func() bool {
			y := rand.Intn(9)
			x := rand.Intn(9)
			cell := board.Get(x, y)
			if cell.IsSet() {
				cell.Set(0)
				return true
			}
			return false
		})
	}

	Generator.SwapDigits = func(board Board) {
		times := rand.Intn(20) + 1
		for i := 0; i < times; i++ {
			a := rand.Intn(9) + 1
			b := rand.Intn(9) + 1
			if a != b {
				swapDigits(board, a, b)
			}
		}
	}

	Generator.SwapRows = func(board Board) {
		times := rand.Intn(20) + 1
		for i := 0; i < times; i++ {
			block := rand.Intn(3) * 3
			a := rand.Intn(3) + block
			b := rand.Intn(3) + block
			if a != b {
				swapRows(board, a, b)
			}
		}
	}

	Generator.SwapColumns = func(board Board) {
		times := rand.Intn(20) + 1
		for i := 0; i < times; i++ {
			block := rand.Intn(3) * 3
			a := rand.Intn(3) + block
			b := rand.Intn(3) + block
			if a != b {
				swapColumns(board, a, b)
			}
		}
	}
}

func swapDigits(board Board, a, b int) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			cell := board.Get(x, y)
			switch cell.Value() {
			case a:
				cell.Reset()
				cell.Set(b)
			case b:
				cell.Reset()
				cell.Set(a)
			default:
			}
		}
	}
}

func swapRows(board Board, a, b int) {
	for x := 0; x < 9; x++ {
		board.Get(x, a).swap(board.Get(x, b))
	}
}

func swapColumns(board Board, a, b int) {
	for y := 0; y < 9; y++ {
		board.Get(a, y).swap(board.Get(b, y))
	}
}
