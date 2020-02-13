package main

type Mediator interface{
	CanArrive(Train) bool
	NotifyFree()
}
