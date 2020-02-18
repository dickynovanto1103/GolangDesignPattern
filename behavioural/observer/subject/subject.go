package subject

import (
	observer2 "GolangDesignPattern/behavioural/observer/observer"
)

type Subject interface {
	AddObserver(observer observer2.Observer)
	RemoveObserver(observer observer2.Observer)
	NotifyObservers()
}
