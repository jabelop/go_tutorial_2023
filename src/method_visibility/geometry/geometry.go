package geometry

type Triangle struct {
	size int
}

/*
	Methods and properties defined with lowercase at the begining like doubleSize and size are considered private,
	they will raise an error if are accessed from another package.
	Methods and properties defined with Uppercase at the begining like SetSize and Perimeter are considered public,
	they can be accessed from another package.
*/
func (t *Triangle) doubleSize() {
	t.size *= 2
}

func (t *Triangle) SetSize(size int) {
	t.size = size
}

func (t *Triangle) Perimeter() int {
	t.doubleSize()
	return t.size * 3
}
