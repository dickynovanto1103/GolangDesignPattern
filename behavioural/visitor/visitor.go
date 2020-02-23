package main

type Visitor interface {
	VisitSquare(s Shape)
	VisitCircle(s Shape)
}
