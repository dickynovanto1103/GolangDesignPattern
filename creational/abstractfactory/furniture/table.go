package furniture

type Table struct{
	name string
}

func NewTable(name string) *Table {
	return &Table{name:name}
}

func (t *Table) GetName() string {
	return t.name
}