package sudoku

import "bytes"

type Board interface {
	Get(int, int) Cell
	Fix()
	String() string
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
	board := &boardType{}
	board.cells = make([]Cell, 0, 81)
	for _, i := range boardReference {
		board.cells = append(board.cells, NewCell(i))
	}
	board.Fix()
	return board
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
