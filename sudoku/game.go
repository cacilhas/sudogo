package sudoku

type Game interface {
	Autofill()
	GameOver() bool
	Get(int, int) Cell
	Set(int, int, int) bool
	Toggle(int, int, int) bool
	Undo() bool
	Redo() bool
	String() string
}

type gameType struct {
	current Board
	undo    []Board
	redo    []Board
}

var emptyBoard Board

func init() {
	emptyBoard = NewBoard()
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			emptyBoard.Get(x, y).Set(0)
		}
	}
	emptyBoard.Fix()
}

func NewGame(input interface{}) (Game, error) {
	var board Board
	switch value := input.(type) {
	case Level:
		board = NewBoard()
		Generator.Generate(board, value)
	case string:
		var err error
		if board, err = LoadBoard(value); err != nil {
			return nil, err
		}
		board.Fix()
	}
	return &gameType{
		current: board,
		undo:    nil,
		redo:    nil,
	}, nil
}

func (game *gameType) Autofill() {
	found := false
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			cell := game.Get(x, y)
			if cell.Candidates() == 1 {
				if !found {
					found = true
					game.addRound()
				}
				for i := 1; i <= 9; i++ {
					if cell.Candidate(i) {
						cell.Set(i)
						break
					}
				}
			}
		}
	}
	if found {
		game.current.partialFix()
	}
}

func (game gameType) GameOver() bool {
	return game.current.GameOver()
}

func (game gameType) Get(x, y int) Cell {
	return game.current.Get(x, y)
}

func (game *gameType) Set(x, y, value int) bool {
	if game.current.Get(x, y).Candidate(value) {
		game.addRound()
		game.current.Get(x, y).Set(value)
		game.current.partialFix()
		return true
	}
	return false
}

func (game *gameType) Toggle(x, y, value int) bool {
	game.addRound()
	// FIXME: the cell has been toggled correctly, but, after the procedure,
	// the board keeps the wrong version of the cell.
	return game.current.Get(x, y).Toggle(value)
}

func (game *gameType) Undo() bool {
	if len(game.undo) == 0 {
		return false
	}
	game.redo = append(game.redo, game.current)
	game.current = game.undo[len(game.undo)-1]
	game.undo = game.undo[:len(game.undo)-1]
	return true
}

func (game *gameType) Redo() bool {
	if len(game.redo) == 0 {
		return false
	}
	game.undo = append(game.undo, game.current)
	game.current = game.redo[len(game.redo)-1]
	game.redo = game.redo[:len(game.redo)-1]
	return true
}

func (game gameType) String() string {
	return game.current.String()
}

func (game *gameType) addRound() {
	game.undo = append(game.undo, game.current)
	game.current = game.current.Clone()
}
