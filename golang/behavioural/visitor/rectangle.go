package main

type rectangle struct {
	l int
	b int
}

func (r *rectangle) accept(v visitor) {
	v.visitForRectangle(r)
}

func (r *rectangle) getType() string {
	return "Rectangle"
}
