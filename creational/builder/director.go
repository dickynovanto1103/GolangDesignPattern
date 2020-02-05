package main

type Director struct {
	builder HPBuilder
}

func NewDirector(builder HPBuilder) *Director {
	return &Director{builder:builder}
}

func (d *Director) BuildHP() *HP {
	d.builder.SetMerk()
	d.builder.SetCountry()
	return d.builder.GetHP()
}
