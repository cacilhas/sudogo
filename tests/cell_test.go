package tests

import (
	"fmt"
	"testing"

	"github.com/cacilhas/sudogo/sudoku"
)

func TestCell(t *testing.T) {
	t.Run("NewCell", func(t *testing.T) {
		for i := 0; i <= 9; i++ {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				if got := sudoku.NewCell(i); got.Value() != i {
					t.Fatalf("expected 0, got %s", got)
				}
			})
		}
		t.Run("negative parameter", func(t *testing.T) {
			if got := sudoku.NewCell(-1); got != nil {
				t.Fatalf("expected nil, got %s", got)
			}
		})
		t.Run("too large parameter", func(t *testing.T) {
			if got := sudoku.NewCell(10); got != nil {
				t.Fatalf("expected nil, got %s", got)
			}
		})
	})
	t.Run("Set", func(t *testing.T) {
		t.Run("0", func(t *testing.T) {
			cell := sudoku.NewCell(2)
			if !cell.Set(0) {
				t.Fatal("could not set cell to 0")
			}
			if cell.IsSet() {
				t.Fatalf("expected empty, got %s", cell)
			}
		})
		for i := 1; i <= 9; i++ {
			t.Run("i", func(t *testing.T) {
				cell := sudoku.NewCell(0)
				if !cell.Set(i) {
					t.Fatalf("could not set cell to %d", i)
				}
				if !cell.IsSet() {
					t.Fatal("expected set, got empty")
				}
				if cell.Value() != i {
					t.Fatalf("expected %d, got %s", i, cell)
				}
			})
		}
		t.Run("negative parameter", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Set(-1) {
				t.Fatal("should not be able to set cell to -1")
			}
			if cell.IsSet() {
				t.Fatalf("expected empty, got %s", cell)
			}
		})
		t.Run("too large parameter", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Set(10) {
				t.Fatal("should not be able to set cell to 10")
			}
			if cell.IsSet() {
				t.Fatalf("expected empty, got %s", cell)
			}
		})
	})
	t.Run("Reset", func(t *testing.T) {
		cell := sudoku.NewCell(2)
		for i := 1; i <= 9; i++ {
			cell.Disable(i)
		}
		if got := cell.Debug(); got != "[2]" {
			t.Fatalf("expected [2], got %s", got)
		}
		cell.Reset()
		if got := cell.Debug(); got != "987654321[2]" {
			t.Fatalf("expected 987654321[2], got %s", got)
		}
	})
	t.Run("Candidate", func(t *testing.T) {
		t.Run("enabled", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			for i := 0; i <= 9; i++ {
				t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
					if !cell.Candidate(i) {
						t.Fatal("expected to be a candidate")
					}
				})
			}
			t.Run("negative parameter", func(t *testing.T) {
				if cell.Candidate(-1) {
					t.Fatal("expected not to be a candidate")
				}
			})
			t.Run("too large parameter", func(t *testing.T) {
				if cell.Candidate(10) {
					t.Fatal("expected not to be a candidate")
				}
			})
		})
		t.Run("disabled", func(t *testing.T) {
			for i := 1; i <= 9; i++ {
				t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
					cell := sudoku.NewCell(0)
					if !cell.Disable(i) {
						t.Fatal("could not disable cell")
					}
					if cell.Candidate(i) {
						t.Fatal("should be disabled")
					}
					if cell.Set(i) {
						t.Fatal("should not be able to set cell")
					}
					if cell.IsSet() {
						t.Fatalf("expected empty, got %s", cell)
					}
				})
			}
		})
	})
	t.Run("Enable", func(t *testing.T) {
		t.Run("negative parameter", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Enable(-1) {
				t.Fatal("should not be able to enable")
			}
		})
		t.Run("too large parameter", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Enable(10) {
				t.Fatal("should not be able to enable")
			}
		})
		t.Run("0", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if !cell.Enable(0) {
				t.Fatal("should always be abled")
			}
		})
	})
	t.Run("Disable", func(t *testing.T) {
		t.Run("negative parameter", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Disable(-1) {
				t.Fatal("should not be able to disable")
			}
		})
		t.Run("too large parameter", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Disable(10) {
				t.Fatal("should not be able to disable")
			}
		})
		t.Run("0", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Disable(0) {
				t.Fatal("should not be able to disable")
			}
		})
		for i := 1; i <= 9; i++ {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				cell := sudoku.NewCell(0)
				if !cell.Disable(i) {
					t.Fatal("could not disable cell")
				}
				if cell.Candidate(i) {
					t.Fatal("should be disabled")
				}
				for j := 1; j <= 9; j++ {
					if j != i {
						if !cell.Candidate(j) {
							t.Fatal("should be a candidate")
						}
					}
				}
				if !cell.Enable(i) {
					t.Fatal("could not enable cell back")
				}
				if !cell.Candidate(i) {
					t.Fatal("should be a candidate")
				}
			})
		}
	})
	t.Run("Toggle", func(t *testing.T) {
		t.Run("negative parameter", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Toggle(-1) {
				t.Fatal("should not be able to disable")
			}
		})
		t.Run("too large parameter", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Toggle(10) {
				t.Fatal("should not be able to disable")
			}
		})
		t.Run("0", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			if cell.Toggle(0) {
				t.Fatal("should not be able to disable")
			}
		})
		for i := 1; i <= 9; i++ {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				cell := sudoku.NewCell(0)
				if !cell.Toggle(i) {
					t.Fatal("could not toggle cell")
				}
				for j := 1; j <= 9; j++ {
					if j == i {
						if cell.Candidate(j) {
							t.Fatalf("[%d] should be disabled", j)
						}
					} else {
						if !cell.Candidate(j) {
							t.Fatalf("[%d] should be a candidate", j)
						}
					}
				}
				if !cell.Toggle(i) {
					t.Fatal("could not toggle cell back")
				}
				for j := 1; j <= 9; j++ {
					if !cell.Candidate(j) {
						t.Fatalf("[%d] should be a candidate", j)
					}
				}
			})
		}
	})
	t.Run("Copy", func(t *testing.T) {
		cell := sudoku.NewCell(3)
		cell.Disable(5)
		got := sudoku.NewCell(0)
		got.Copy(cell)
		cell.Set(0)
		cell.Enable(5)
		cell.Disable(3)
		if got.Value() != 3 {
			t.Fatalf("expected 3, got %s", got)
		}
		if got.Candidate(5) {
			t.Fatalf("should not be a candidate: %s", got.Debug())
		}
		if !got.Candidate(3) {
			t.Fatalf("should be a candidate: %s", got.Debug())
		}
	})
	t.Run("Debug", func(t *testing.T) {
		t.Run("0", func(t *testing.T) {
			cell := sudoku.NewCell(0)
			expected := "987654321[.]"
			if got := cell.Debug(); got != expected {
				t.Fatalf("expected %s, got %s", expected, got)
			}
		})
		for i := 1; i <= 9; i++ {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				cell := sudoku.NewCell(i)
				expected := fmt.Sprintf("987654321[%d]", i)
				if got := cell.Debug(); got != expected {
					t.Fatalf("expected %s, got %s", expected, got)
				}
			})
		}
		for j := 1; j <= 9; j++ {
			t.Run(fmt.Sprintf("0-%d", 0), func(t *testing.T) {
				cell := sudoku.NewCell(0)
				cell.Disable(j)
				expected := ""
				for i := 9; i >= 1; i-- {
					if i != j {
						expected = fmt.Sprintf("%s%d", expected, i)
					}
				}
				expected += "[.]"
				if got := cell.Debug(); got != expected {
					t.Fatalf("expected %s, got %s", expected, got)
				}
			})
			for i := 1; i <= 9; i++ {
				t.Run(fmt.Sprintf("%d-%d", i, j), func(t *testing.T) {
					cell := sudoku.NewCell(i)
					cell.Disable(j)
					expected := ""
					for ii := 9; ii >= 1; ii-- {
						if ii != j {
							expected = fmt.Sprintf("%s%d", expected, ii)
						}
					}
					expected = fmt.Sprintf("%s[%d]", expected, i)
					if got := cell.Debug(); got != expected {
						t.Fatalf("expected %s, got %s", expected, got)
					}
				})
			}
		}
	})
}
