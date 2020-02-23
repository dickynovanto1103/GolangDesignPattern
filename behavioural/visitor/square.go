package main

type Square struct {
	side int
}

func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s)
}
