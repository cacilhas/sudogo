package tests

import (
	"testing"

	"github.com/cacilhas/sudogo/sudoku"
)

func TestLevel(t *testing.T) {
	t.Run("Min", func(t *testing.T) {
		t.Run("EXTREMELY_EASY", func(t *testing.T) {
			if got := sudoku.EXTREMELY_EASY.Min(); got != 25 {
				t.Fatalf("expected 25, got %d", got)
			}
		})
		t.Run("EASY", func(t *testing.T) {
			if got := sudoku.EASY.Min(); got != 32 {
				t.Fatalf("expected 32, got %d", got)
			}
		})
		t.Run("MEDIUM", func(t *testing.T) {
			if got := sudoku.MEDIUM.Min(); got != 45 {
				t.Fatalf("expected 45, got %d", got)
			}
		})
		t.Run("HARD", func(t *testing.T) {
			if got := sudoku.HARD.Min(); got != 50 {
				t.Fatalf("expected 50, got %d", got)
			}
		})
		t.Run("FIENDISH", func(t *testing.T) {
			if got := sudoku.FIENDISH.Min(); got != 54 {
				t.Fatalf("expected 54, got %d", got)
			}
		})
	})
	t.Run("Max", func(t *testing.T) {
		t.Run("EXTREMELY_EASY", func(t *testing.T) {
			if got := sudoku.EXTREMELY_EASY.Max(); got != 31 {
				t.Fatalf("expected 31, got %d", got)
			}
		})
		t.Run("EASY", func(t *testing.T) {
			if got := sudoku.EASY.Max(); got != 44 {
				t.Fatalf("expected 44, got %d", got)
			}
		})
		t.Run("MEDIUM", func(t *testing.T) {
			if got := sudoku.MEDIUM.Max(); got != 49 {
				t.Fatalf("expected 49, got %d", got)
			}
		})
		t.Run("HARD", func(t *testing.T) {
			if got := sudoku.HARD.Max(); got != 53 {
				t.Fatalf("expected 53, got %d", got)
			}
		})
		t.Run("FIENDISH", func(t *testing.T) {
			if got := sudoku.FIENDISH.Max(); got != 59 {
				t.Fatalf("expected 59, got %d", got)
			}
		})
		t.Run("Exec", func(t *testing.T) {
			t.Skip("TODO: how to test this?")
		})
	})
}
