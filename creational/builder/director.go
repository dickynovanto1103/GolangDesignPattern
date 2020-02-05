package main

type Director struct {
	builder HPBuilder
}

func NewDirector(builder HPBuilder) *Director {
	return &Director{builder:builder}
}

func (d *Director) Build() HPBuilder {
	d.builder.SetMerk()
	d.builder.SetCountry()
	return d.builder
}
