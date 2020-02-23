package main

type Shape interface {
	Accept(v Visitor)
}
