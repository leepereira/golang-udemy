package shapes

type Rectangle struct {
	Length  int
	Breadth int
}

func (r Rectangle) Area() int {
	return r.Length * r.Breadth
}
