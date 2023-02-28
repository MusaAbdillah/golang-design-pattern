package main

type Rectangle struct {
	l int
	b int
}

func (s *Rectangle) getType() string {
	return "Rectangle"
}

func (s *Rectangle) accept(v Visitor) {
	v.visitForRectangle(s)
}
