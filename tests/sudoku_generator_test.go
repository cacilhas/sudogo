package tests

import (
	"math/rand"
	"testing"

	"github.com/cacilhas/sudogo/sudoku"
)

func TestSokuGenerator(t *testing.T) {
	t.Run("Board", func(t *testing.T) {
		t.Run("SOLVED", func(t *testing.T) {
			rand.Seed(1)
			board := sudoku.NewBoard(sudoku.SOLVED)
			expected := " 694 218 573\n 218 573 694\n 573 694 218\n\n"
			expected += " 942 185 736\n 185 736 942\n 736 942 185\n\n"
			expected += " 421 857 369\n 857 369 421\n 369 421 857\n"
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%v\ngot:\n%v", expected, got)
			}
		})
		t.Run("IsDone", func(t *testing.T) {
			for i := 0; i < 10; i++ {
				rand.Seed(int64(i))
				board := sudoku.NewBoard(sudoku.SOLVED)
				if !board.IsDone() {
					t.Fatalf("expected board to be solved:\n%v", board)
				}
			}
		})
		t.Run("EASY", func(t *testing.T) {
			rand.Seed(2)
			board := sudoku.NewBoard(sudoku.EASY)
			expected := " 8.5 6.. 273\n .9. 2.3 ..5\n 27. ... ..1\n\n"
			expected += " 456 9.2 73.\n 912 7.8 4.6\n 7.. .56 912\n\n"
			expected += " ... .27 .8.\n ... ..4 ..9\n .8. ... .27\n"
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%v\ngot:\n%v", expected, got)
			}
		})
		t.Run("MEDIUM", func(t *testing.T) {
			rand.Seed(3)
			board := sudoku.NewBoard(sudoku.MEDIUM)
			expected := " ..1 .7. ...\n ... ... ...\n ... .3. .79\n\n"
			expected += " 3.4 792 6.5\n 79. ... ...\n ... .14 ..2\n\n"
			expected += " 1.7 926 .53\n 9.. .5. ...\n 8.. .4. 926\n"
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%v\ngot:\n%v", expected, got)
			}
		})
		t.Run("HARD", func(t *testing.T) {
			rand.Seed(4)
			board := sudoku.NewBoard(sudoku.HARD)
			expected := " ... ... ..2\n 9.. ... .64\n ..2 8.. ...\n\n"
			expected += " ... .51 ...\n ... ..8 ..9\n 32. 649 7..\n\n"
			expected += " ..7 513 ...\n ... ..6 ...\n ... ... ...\n"
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%v\ngot:\n%v", expected, got)
			}
		})
		t.Run("FIENDISH", func(t *testing.T) {
			rand.Seed(5)
			board := sudoku.NewBoard(sudoku.FIENDISH)
			expected := " ... .7. 2..\n ..1 .9. ...\n ... ... ...\n\n"
			expected += " ... ... ...\n .12 ... ..3\n 9.. ... 7.2\n\n"
			expected += " 8.. ... ...\n ... ... ...\n ... ... ...\n"
			if got := board.String(); got != expected {
				t.Fatalf("expected:\n%v\ngot:\n%v", expected, got)
			}
		})
	})
}
