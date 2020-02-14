package main

type Originator struct {
	state string
}

func (o *Originator) getState() string {
	return o.state
}

func (o *Originator) setState(s string) {
	o.state = s
}

func (o *Originator) createMemento() *memento {
	return &memento{state: o.getState()}
}

func (o *Originator) restoreMemento(m *memento) string {
	return m.getSavedState()
}