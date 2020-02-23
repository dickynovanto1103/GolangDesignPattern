package main

type Circle struct {
	r int
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}
