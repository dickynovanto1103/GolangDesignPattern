package main

type Original struct {
	state string
}

func (o *Original) getState() string {
	return o.state
}

func (o *Original) setState(s string) {
	o.state = s
}

func (o *Original) createMemento() *memento {
	return &memento{state: o.getState()}
}

func (o *Original) restoreMemento(m *memento) string {
	return m.getSavedState()
}