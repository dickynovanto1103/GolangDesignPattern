package furniture

type Chair struct{
	name string
}

func NewChair(name string) *Chair {
	return &Chair{name:name}
}

func (c *Chair) GetName() string {
	return c.name
}
