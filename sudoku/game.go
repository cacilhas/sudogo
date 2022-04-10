package sudoku

type Game interface {
	Get(int, int) Cell
	Set(int, int, int) bool
	Toggle(int, int, int) bool
	Undo() bool
	Redo() bool
	String() string
}

type gameType struct {
	rounds []Board
	index  int
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

func NewGame(level Level) Game {
	rounds := make([]Board, 0, 1)
	board := NewBoard()
	Generator.Generate(board, level)
	return &gameType{
		rounds: append(rounds, board),
		index:  0,
	}
}

func (game gameType) Get(x, y int) Cell {
	board := game.current()
	if board == nil {
		return nil
	}
	return board.Get(x, y)
}

func (game *gameType) Set(x, y, value int) bool {
	board := game.current()
	if board == nil {
		return false
	}
	if board.Get(x, y).Candidate(value) {
		game.addRound()
		board = game.current()
		board.Get(x, y).Set(value)
		board.Fix()
		return true
	}
	return false
}

func (game *gameType) Toggle(x, y, value int) bool {
	if game.current() == nil {
		return false
	}
	game.addRound()
	// FIXME: the cell has been toggled correctly, but, after the procedure,
	// the board keeps the wrong version of the cell.
	return game.current().Get(x, y).Toggle(value)
}

func (game *gameType) Undo() bool {
	if game.index <= 0 {
		return false
	}
	game.index--
	return true
}

func (game *gameType) Redo() bool {
	if game.index == len(game.rounds)-1 {
		return false
	}
	game.index++
	return true
}

func (game gameType) String() string {
	board := game.current()
	if board == nil {
		return emptyBoard.String()
	}
	return board.String()
}

func (game *gameType) addRound() {
	var round Board
	if game.index < 0 {
		round = NewBoard()
	} else {
		round = game.current().Clone()
	}
	game.index++
	game.rounds = append(game.rounds[:game.index], round)
}

func (game *gameType) current() Board {
	if game.index < 0 {
		return nil
	}
	return game.rounds[game.index]
}
