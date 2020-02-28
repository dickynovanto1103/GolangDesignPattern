package main

import "fmt"

type SMSNotifDecorator struct {
	Notifier
}

func NewSMSNotifDecorator(baseNotifDecorator Notifier) *SMSNotifDecorator {
	return &SMSNotifDecorator{Notifier: baseNotifDecorator}
}

func (n *SMSNotifDecorator) Notify() {
	fmt.Println("sms notification")
	n.notifier.Notify()
}
