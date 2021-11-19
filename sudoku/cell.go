package sudoku

type Cell interface {
	Value() uint8
	SetValue(value uint8)
	Clear(uint8)
	Get(uint8) bool
	Set(uint8)
	Toggle(uint8)
}

type cell struct {
	value uint8
	guess [9]bool
}

func NewCell() Cell {
	return &cell{
		value: 0,
		guess: [9]bool{true, true, true, true, true, true, true, true, true},
	}
}

func (c *cell) reset() {
	found := uint8(0)
	for i := 0; i < 9; i++ {
		if c.guess[i] {
			if found != 0 {
				found = 0
				break
			}
			found = uint8(i) + 1
		}
	}
	c.value = found
}

func (c cell) Value() uint8 {
	return c.value
}

func (c *cell) SetValue(v uint8) {
	checkValue(v, true)
	c.value = v
}

func (c *cell) Set(v uint8) {
	checkValue(v, false)
	c.guess[v-1] = true
	c.reset()
}

func (c *cell) Clear(v uint8) {
	checkValue(v, false)
	c.guess[v-1] = false
	c.reset()
}

func (c *cell) Toggle(v uint8) {
	checkValue(v, false)
	c.guess[v-1] = !c.guess[v-1]
}

func (c *cell) Get(v uint8) bool {
	checkValue(v, false)
	return c.guess[v-1]
}
