package tests

import (
	"testing"

	"github.com/cacilhas/sudogo/sudoku"
)

func TestCell(t *testing.T) {
	t.Run("Cell", func(t *testing.T) {
		t.Run("NewCell", func(t *testing.T) {
			cell := sudoku.NewCell()
			if got := cell.Value(); got != 0 {
				t.Fatalf("expected cell to be 0, got %v", got)
			}
			for i := uint8(1); i < 10; i++ {
				if !cell.Get(i) {
					t.Fatalf("expected to be set at %v", i)
				}
			}
		})
		t.Run("Clear and Set", func(t *testing.T) {
			cell := sudoku.NewCell()
			cell.Clear(2)
			for i := uint8(1); i < 10; i++ {
				got := cell.Get(i)
				if i == 2 && got {
					t.Fatalf("expected to be clear at %v", i)
				} else if !(i == 2 || got) {
					t.Fatalf("expected to be set at %v", i)
				}
			}
			cell.Set(2)
			if !cell.Get(2) {
				t.Fatalf("expected to be set at 2")
			}
		})
		t.Run("Toggle", func(t *testing.T) {
			cell := sudoku.NewCell()
			cell.Toggle(2)
			for i := uint8(1); i < 10; i++ {
				got := cell.Get(i)
				if i == 2 && got {
					t.Fatalf("expected to be clear at %v", i)
				} else if !(i == 2 || got) {
					t.Fatalf("expected to be set at %v", i)
				}
			}
			cell.Toggle(2)
			if !cell.Get(2) {
				t.Fatalf("expected to be set at 2")
			}
		})
		t.Run("autoset", func(t *testing.T) {
			cell := sudoku.NewCell()
			for i := uint8(1); i < 9; i++ {
				cell.Clear(i)
			}
			if got := cell.Value(); got != 9 {
				t.Fatalf("expected value to be 9, got %v", got)
			}
			cell.Clear(9)
			cell.Set(5)
			if got := cell.Value(); got != 5 {
				t.Fatalf("expected value to be 5, got %v", got)
			}
			cell.Set(2)
			if got := cell.Value(); got != 0 {
				t.Fatalf("expected value to be 0, got %v", got)
			}
		})
	})
}
