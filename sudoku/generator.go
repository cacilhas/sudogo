package sudoku

import (
	"fmt"
	"math/rand"
)

type Hardship uint8

const (
	NONE     = 0
	SOLVED   = 1
	EASY     = 40
	MEDIUM   = 50
	HARD     = 60
	FIENDISH = 70
)

type Board interface {
	Get(int, int) uint8
	Set(int, int, uint8)
	IsDone() bool
	IsValid(int, int, uint8) bool
	String() string
}

type board struct {
	data [81]uint8
}

func NewBoard(hardship Hardship) Board {
	var res board
	init_value := uint8(1)
	for y := 0; y < 9; y++ {
		value := init_value
		for x := 0; x < 9; x++ {
			res.Set(x, y, value)
			value++
			if value > 9 {
				value = 1
			}
		}
		init_value += 3
		if init_value > 9 {
			init_value -= 8
		}
	}
	switch hardship {
	case NONE:
		// nothing to do
	case SOLVED:
		res.shuffle()
	default:
		res.shuffle()
		res.hide(int(hardship))
	}
	return &res
}

func checkVec2(x, y int) {
	if x < 0 || x > 8 || y < 0 || y > 8 {
		panic("out of range")
	}
}

func checkValue(v uint8) {
	if v > 9 {
		panic("invalid value")
	}
}

func (b *board) Get(x, y int) uint8 {
	checkVec2(x, y)
	return b.data[x+y*9]
}

func (b *board) Set(x, y int, v uint8) {
	checkVec2(x, y)
	checkValue(v)
	b.data[x+y*9] = v
}

func (b board) IsValid(x, y int, v uint8) bool {
	if v > 9 {
		return false
	}
	if v == 0 {
		return true
	}

	for lx := 0; lx < 9; lx++ {
		if lx != x && b.Get(lx, y) == v {
			return false
		}
	}
	for ly := 0; ly < 9; ly++ {
		if ly != y && b.Get(x, ly) == v {
			return false
		}
	}
	bx := int(x/3) * 3
	by := int(y/3) * 3

	for ly := by; ly < by+3; ly++ {
		for lx := bx; lx < bx+3; lx++ {
			if lx != x && ly != y && b.Get(lx, ly) == v {
				return false
			}
		}
	}

	return true
}

func (b board) IsDone() bool {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			v := b.Get(x, y)
			if v == 0 && !b.IsValid(x, y, v) {
				return false
			}
		}
	}
	return true
}

func (b board) String() string {
	buf := ""
	for y := 0; y < 9; y++ {
		if y > 0 && y%3 == 0 {
			buf += "\n"
		}
		for x := 0; x < 9; x++ {
			if x%3 == 0 {
				buf += " "
			}
			cur := b.Get(x, y)
			if cur == 0 {
				buf += "."
			} else {
				buf = fmt.Sprintf("%s%d", buf, cur)
			}
		}
		buf += "\n"
	}
	return buf
}

func (b *board) shuffle() {
	values := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 9; i++ {
		j := rand.Intn(9-i) + i
		values[i], values[j] = values[j], values[i]
	}
	for i := 0; i < 81; i++ {
		b.data[i] = values[b.data[i]-1]
	}
}

func findNext(data [81]uint8, from int) int {
	i := from
	for {
		if data[i] != 0 {
			return i
		}
		i = (i + 1) % 81
		if i == from {
			return -1
		}
	}
}

func (b *board) hide(times int) {
	for i := 0; i < times; i++ {
		j := findNext(b.data, rand.Intn(81))
		if j != -1 {
			b.data[j] = 0
		}
	}
}
