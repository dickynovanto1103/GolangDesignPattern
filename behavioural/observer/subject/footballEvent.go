package subject

import (
	observer2 "GolangDesignPattern/behavioural/observer/observer"
)

type FootballEvent struct {
	id string
	observers []observer2.Observer
}

func NewFootballEvent(id string) *FootballEvent {
	return &FootballEvent{
		id:        id,
	}
}

func (e *FootballEvent) AddObserver(observer observer2.Observer) {
	e.observers = append(e.observers, observer)
}

func (e *FootballEvent) RemoveObserver(observer observer2.Observer) {
	lenObservers := len(e.observers)
	for idx, ob := range e.observers {
		if ob == observer {
			e.observers[idx] = e.observers[lenObservers-1]
			break
		}
	}
	e.observers = e.observers[:lenObservers-1]
}

func (e *FootballEvent) NotifyObservers() {
	for _, ob := range e.observers {
		ob.Update(e.id)
	}
}