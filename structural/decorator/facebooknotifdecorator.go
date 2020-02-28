package main

import "fmt"

type FacebookNotifDecorator struct {
	Notifier
}

func NewFacebookNotifDecorator(baseNotifDecorator Notifier) Notifier {
	return &FacebookNotifDecorator{Notifier: baseNotifDecorator}
}

func (f *FacebookNotifDecorator) Notify() {
	fmt.Println("facebook notification")
	f.notifier.Notify()
}
