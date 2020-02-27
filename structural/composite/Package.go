package main

type Package struct {
	children []Component
}

func NewPackage() *Package {
	return &Package{}
}

func (p *Package) Count() int {
	sum := 0
	for _, c := range p.children {
		sum += c.Count()
	}
	return sum
}

func (p *Package) AddComponent(c Component) {
	p.children = append(p.children, c)
}
