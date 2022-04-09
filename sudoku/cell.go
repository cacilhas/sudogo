package sudoku

import (
	"bytes"
	"fmt"
)

type Cell interface {
	Value() int
	IsSet() bool
	Set(int) bool
	Reset()
	Candidate(int) bool
	Enable(int) bool
	Disable(int) bool
	Toggle(int) bool
	Clone() Cell
	Copy(Cell)
	String() string
	Debug() string
	realValue() uint16
}

type cellType struct {
	uint16
}

func NewCell(value int) Cell {
	if value < 0 || value > 9 {
		return nil
	}
	return &cellType{uint16(0x1ff0 | value)}
}

func (cell cellType) Value() int {
	return int(cell.uint16 & 0xf)
}

func (cell cellType) IsSet() bool {
	return cell.uint16&0xf != 0
}

func (cell *cellType) Set(value int) bool {
	if value < 0 || value > 9 {
		return false
	}
	if !cell.Candidate(value) {
		return false
	}
	cell.uint16 = uint16(0x1ff0 | value)
	return true
}

func (cell *cellType) Reset() {
	cell.uint16 = uint16(0x1ff0 | (cell.uint16 & 0xf))
}

func (cell cellType) Candidate(index int) bool {
	if index < 0 || index > 9 {
		return false
	}
	if index == 0 {
		return true
	}
	var modifier uint16 = 0x8 << index
	return cell.uint16&modifier != 0
}

func (cell *cellType) Enable(index int) bool {
	if index < 0 || index > 9 {
		return false
	}
	if index == 0 {
		return true
	}
	var modifier uint16 = 0x8 << index
	cell.uint16 |= modifier
	return true
}

func (cell *cellType) Disable(index int) bool {
	if index <= 0 || index > 9 {
		return false
	}
	var modifier uint16 = 0x8 << index
	cell.uint16 &= ^modifier
	return true
}

func (cell *cellType) Toggle(index int) bool {
	if index <= 0 || index > 9 {
		return false
	}
	var modifier uint16 = 0x8 << index
	cell.uint16 ^= modifier
	return true
}

func (cell cellType) Clone() Cell {
	return &cellType{cell.uint16}
}

func (cell *cellType) Copy(other Cell) {
	cell.uint16 = other.realValue()
}

func (cell cellType) String() string {
	if cell.IsSet() {
		return fmt.Sprintf("%d", cell.uint16&0xf)
	}
	return "."
}

func (cell cellType) Debug() string {
	var buf bytes.Buffer
	for i := 9; i > 0; i-- {
		if cell.Candidate(i) {
			buf.WriteByte('0' + byte(i))
		}
	}
	return fmt.Sprintf("%s[%s]", buf.String(), cell)
}

func (cell cellType) realValue() uint16 {
	return cell.uint16
}
