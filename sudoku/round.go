package sudoku

type Round interface {
	Board() Board
	Get(int, int) Cell
	Set(int, int, interface{})
}

type round struct {
	source Board
	state  [81]Cell
}

func NewRound(board Board) Round {
	r := round{source: board}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if v := board.Get(x, y); v != 0 {
				r.Set(x, y, v)
			}
		}
	}
	return &r
}

func (r round) Board() Board {
	board := newEmptyBoard()
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			board.Set(x, y, r.Get(x, y).Value())
		}
	}
	return board
}

func (r round) Get(x, y int) Cell {
	checkVec2(x, y)
	return r.state[x+y*9]
}

func (r *round) Set(x, y int, v interface{}) {
	checkVec2(x, y)

	switch vv := v.(type) {
	case uint8:
		r.state[x+y*9].SetValue(vv)

	case Cell:
		r.state[x+y*9] = vv
	}

	if value := r.Get(x, y).Value(); value != 0 {
		for lx := 0; x < 9; x++ { // clean up columns
			if lx != x {
				r.state[lx+y*9].Clear(value)
			}
		}
		for ly := 0; y < 9; y++ { // clean up rows
			if ly != y {
				r.state[x+ly*9].Clear(value)
			}
		}
		bx := int(x/3) * 3
		by := int(y/3) * 3
		for ly := by; ly < by+2; ly++ { // clean up blocks
			for lx := bx; lx < bx+2; lx++ {
				if lx != x || ly != y {
					r.state[lx+ly*9].Clear(value)
				}
			}
		}
	}
}
