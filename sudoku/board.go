package sudoku

import (
	"bytes"
	"fmt"
)

type Board interface {
	Get(int, int) Cell
	GameOver() bool
	Fix()
	String() string
	Clone() Board
	partialFix()
}

type boardType struct {
	cells []Cell
}

var boardReference [81]int = [81]int{
	1, 2, 3, 4, 5, 6, 7, 8, 9,
	4, 5, 6, 7, 8, 9, 1, 2, 3,
	7, 8, 9, 1, 2, 3, 4, 5, 6,
	2, 3, 4, 5, 6, 7, 8, 9, 1,
	5, 6, 7, 8, 9, 1, 2, 3, 4,
	8, 9, 1, 2, 3, 4, 5, 6, 7,
	3, 4, 5, 6, 7, 8, 9, 1, 2,
	6, 7, 8, 9, 1, 2, 3, 4, 5,
	9, 1, 2, 3, 4, 5, 6, 7, 8,
}

func NewBoard() Board {
	board := &boardType{cells: make([]Cell, 0, 81)}
	for _, i := range boardReference {
		board.cells = append(board.cells, NewCell(i))
	}
	board.Fix()
	return board
}

func (board boardType) GameOver() bool {
	for _, cell := range board.cells {
		if !cell.IsSet() {
			return false
		}
	}
	return true
}

func LoadBoard(s string) (Board, error) {
	board := &boardType{cells: make([]Cell, 0, 81)}
	for _, c := range s {
		if c == '.' {
			board.cells = append(board.cells, NewCell(0))
		} else if c >= '0' && c <= '9' {
			board.cells = append(board.cells, NewCell(int(c-'0')))
		}
	}
	if len(board.cells) == 81 {
		return board, nil
	}
	return nil, fmt.Errorf("wrong board size, expected 81, got %d", len(board.cells))
}

func (board *boardType) Get(x, y int) Cell {
	if x < 0 || x >= 9 || y < 0 || y >= 9 {
		return nil
	}
	return board.cells[y*9+x]
}

func (board *boardType) Fix() {
	for _, cell := range board.cells {
		cell.Reset()
	}
	board.partialFix()
}

func (board *boardType) partialFix() {
	for y := 0; y < 9; y++ {
		fixRange(board.row(y))
	}
	for x := 0; x < 9; x++ {
		fixRange(board.column(x))
	}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fixRange(board.block(x, y))
		}
	}
}

func (board boardType) Clone() Board {
	res := &boardType{cells: make([]Cell, 0, 81)}
	for i := 0; i < 81; i++ {
		res.cells = append(res.cells, board.cells[i].Clone())
	}
	return res
}

func (board boardType) String() string {
	var buf bytes.Buffer
	for y := 0; y < 9; y++ {
		if y%3 == 0 {
			buf.WriteString("+---+---+---+\n")

		}
		for x := 0; x < 9; x++ {
			if x%3 == 0 {
				buf.WriteByte('|')
			}
			buf.WriteString(board.Get(x, y).String())
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

func fixRange(cells []Cell) {
	for _, cell := range cells {
		if cell.IsSet() {
			value := cell.Value()
			for _, other := range cells {
				if other.Value() != value {
					other.Disable(value)
				}
			}
		}
	}
}

func (board *boardType) row(y int) []Cell {
	return board.cells[y*9 : y*9+9]
}

func (board *boardType) column(x int) []Cell {
	res := make([]Cell, 0, 9)
	for y := 0; y < 9; y++ {
		res = append(res, board.cells[y*9+x])
	}
	return res
}

func (board *boardType) block(x, y int) []Cell {
	res := make([]Cell, 0, 9)
	for ly := y * 3; ly < y*3+3; ly++ {
		for lx := x * 3; lx < x*3+3; lx++ {
			res = append(res, board.cells[ly*9+lx])
		}
	}
	return res
}
