package sudoku

func checkVec2(x, y int) {
	if x < 0 || x > 8 || y < 0 || y > 8 {
		panic("out of bounds")
	}
}

func checkValue(v uint8, zero bool) {
	if (v == 0 && !zero) || v > 9 {
		panic("invalid value")
	}
}
