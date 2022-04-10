package tests

import (
	"math/rand"
	"testing"

	"github.com/cacilhas/sudogo/sudoku"
)

func TestGame(t *testing.T) {
	t.Run("NewGame", func(t *testing.T) {
		rand.Seed(-1)
		game := sudoku.NewGame(sudoku.EASY)
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
		if got := game.String(); got != expected {
			t.Fatalf("expected:\n%s\ngot:\n%s", expected, got)
		}
	})

	t.Run("Set", func(t *testing.T) {
		t.Run("valid move", func(t *testing.T) {
			rand.Seed(-1)
			game := sudoku.NewGame(sudoku.EASY)
			control := game.String()
			if !game.Get(1, 3).Candidate(3) {
				t.Fatal("expected 3 to be a candidate")
			}
			if !game.Set(0, 3, 3) {
				t.Fatalf("unable to set 0,3 to 3: %s", game.Get(0, 3).Debug())
			}
			if got := game.Get(0, 3); got.Value() != 3 {
				t.Fatalf("expected 3, got %s", got.Debug())
			}
			if game.Get(1, 3).Candidate(3) {
				t.Fatalf("expected 3 not to be a candidate anymore")
			}
			if got := game.String(); got == control {
				t.Fatalf("expected game to be changed:\n%s", got)
			}
		})
		t.Run("invalid move", func(t *testing.T) {
			rand.Seed(-1)
			game := sudoku.NewGame(sudoku.EASY)
			control := game.String()
			if game.Set(0, 3, 1) {
				t.Fatal("expected Set to fail")
			}
			if got := game.String(); got != control {
				t.Fatalf("expected:\n%s\ngot:\n%s", control, got)
			}
		})
	})

	t.Run("Undo", func(t *testing.T) {
		t.Run("no undo/redo", func(t *testing.T) {
			rand.Seed(-1)
			game := sudoku.NewGame(sudoku.EASY)
			control := game.String()
			if game.Undo() {
				t.Fatal("expected Undo to fail")
			}
			if got := game.String(); got != control {
				t.Fatalf("expected:\n%s\ngot:\n%s", control, got)
			}
			if game.Redo() {
				t.Fatal("expected Redo to fail")
			}
			if got := game.String(); got != control {
				t.Fatalf("expected:\n%s\ngot:\n%s", control, got)
			}
		})
		t.Run("undo/redo", func(t *testing.T) {
			rand.Seed(-1)
			game := sudoku.NewGame(sudoku.EASY)
			control := game.String()
			game.Set(0, 3, 3)
			if got := game.String(); got == control {
				t.Fatalf("expected game to be changed:\n%s", got)
			}
			if !game.Undo() {
				t.Fatal("expected Undo to succeed")
			}
			if got := game.String(); got != control {
				t.Fatalf("expected:\n%s\ngot:\n%s", control, got)
			}
			if !game.Redo() {
				t.Fatal("expected Redo to succeed")
			}
			if got := game.String(); got == control {
				t.Fatalf("expected game to be changed:\n%s", got)
			}
		})
	})

	t.Run("Toggle", func(t *testing.T) {
		t.Run("disable", func(t *testing.T) {
			rand.Seed(-1)
			game := sudoku.NewGame(sudoku.EASY)
			if !game.Get(0, 3).Candidate(3) {
				t.Fatalf("expected 3 to be a candidate: %s", game.Get(1, 3).Debug())
			}
			if !game.Toggle(0, 3, 3) {
				t.Fatal("expected Toggle to succeed")
			}
			if game.Get(0, 3).Candidate(3) {
				t.Fatalf("expected 3 not to be a candidate anymore: %s", game.Get(1, 3).Debug())
			}
			if !game.Undo() {
				t.Fatal("expected Undo to succeed")
			}
			if !game.Get(0, 3).Candidate(3) {
				t.Fatal("expected 3 to be a candidate again")
			}
		})
		t.Run("enable", func(t *testing.T) {
			rand.Seed(-1)
			game := sudoku.NewGame(sudoku.EASY)
			if game.Get(0, 3).Candidate(4) {
				t.Fatalf("expected 4 not to be a candidate: %s", game.Get(1, 3).Debug())
			}
			if !game.Toggle(0, 3, 4) {
				t.Fatal("expected Toggle to succeed")
			}
			if !game.Get(0, 3).Candidate(4) {
				t.Fatalf("expected 4 to be a candidate now: %s", game.Get(1, 3).Debug())
			}
			if !game.Undo() {
				t.Fatal("expected Undo to succeed")
			}
			if game.Get(0, 3).Candidate(4) {
				t.Fatal("expected 3 not to be a candidate anymore")
			}
		})
	})
}
