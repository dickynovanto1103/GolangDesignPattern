package main

type caretaker struct {
	memList []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.memList = append(c.memList, m)
}

func (c *caretaker) getMemento(idx int) *memento {
	return c.memList[idx]
}