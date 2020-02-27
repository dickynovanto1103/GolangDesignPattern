package main

type Object struct {
	price int
}

func NewObject(price int) *Object {
	return &Object{price:price}
}

func (o *Object) Count() int {
	return o.price
}
