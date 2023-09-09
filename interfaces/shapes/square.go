package shapes

type Square struct {
	Side int
}

// Each shape (Square and Rectangle) should implement the Area interface by providing an Area() method.
func (s Square) Area() int {
	return s.Side * s.Side
}
