package sudoku

import (
	"fmt"
	"math/rand"
)

type Level uint16

const (
	EXTREMELY_EASY Level = Level((25 << 8) | 31)
	EASY           Level = Level((32 << 8) | 44)
	MEDIUM         Level = Level((45 << 8) | 49)
	HARD           Level = Level((50 << 8) | 53)
	FIENDISH       Level = Level((54 << 8) | 59)
)

func (level Level) Min() int {
	return int(level >> 8)
}

func (level Level) Max() int {
	return int(level & 0xff)
}

func (level Level) Exec(callback func() bool) {
	min := level.Min()
	max := level.Max()
	for i := 0; i < min; i++ {
		for !callback() {
		}
	}
	stop := rand.Int()
	for i := min; i < max; i++ {
		if stop%(max-i+1) == 0 || !callback() {
			break
		}
	}
}

func (level Level) String() string {
	switch level {
	case EXTREMELY_EASY:
		return "Extremely Easy"
	case EASY:
		return "Easy"
	case MEDIUM:
		return "Medium"
	case HARD:
		return "Hard"
	case FIENDISH:
		return "Fiendish"
	default:
		return fmt.Sprintf("%04X", uint16(level))
	}
}
