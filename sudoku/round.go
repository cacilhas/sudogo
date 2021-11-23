package sudoku

type Round interface {
	Board() Board
	Get(int, int) Cell
	Set(int, int, interface{})
	Undo() bool
	state() [81]Cell
}

type round struct {
	source  Board
	current int
	states  [][81]Cell
}

func NewRound(board Board) Round {
	r := round{source: board}
	var state [81]Cell
	r.states = append(r.states, state)
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if v := board.Get(x, y); v != 0 {
				r.Set(x, y, v)
			}
		}
	}
	return &r
}

func (r *round) state() [81]Cell {
	return r.states[r.current]
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
	return r.state()[x+y*9]
}

func (r *round) Set(x, y int, v interface{}) {
	checkVec2(x, y)
	var state [81]Cell
	for i, cell := range r.state() {
		state[i] = cell.Copy()
	}

	switch vv := v.(type) {
	case uint8:
		state[x+y*9].SetValue(vv)

	case Cell:
		state[x+y*9] = vv
	}

	if value := r.Get(x, y).Value(); value != 0 {
		for lx := 0; x < 9; x++ { // clean up columns
			if lx != x {
				state[lx+y*9].Clear(value)
			}
		}
		for ly := 0; y < 9; y++ { // clean up rows
			if ly != y {
				state[x+ly*9].Clear(value)
			}
		}
		bx := int(x/3) * 3
		by := int(y/3) * 3
		for ly := by; ly < by+2; ly++ { // clean up blocks
			for lx := bx; lx < bx+2; lx++ {
				if lx != x || ly != y {
					state[lx+ly*9].Clear(value)
				}
			}
		}
	}

	r.current++
	if len(state) > r.current {
		r.states[r.current] = state
	} else {
		r.states = append(r.states, state)
	}
}

func (r *round) Undo() bool {
	if r.current == 0 {
		return false
	}
	r.current--
	return true
}
