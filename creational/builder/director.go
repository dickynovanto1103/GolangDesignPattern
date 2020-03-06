package main

type Director struct {
	builder HPBuilder
}

func NewDirector(builder HPBuilder) *Director {
	return &Director{builder:builder}
}

func (d *Director) BuildHP() {
	d.builder.SetMerk()
	d.builder.SetCountry()
}
