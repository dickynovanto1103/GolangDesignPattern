package main

import "fmt"

type BaseNotifDecorator struct {
	notifier Notifier
}

func (n *BaseNotifDecorator) Notify() {
	fmt.Println("base notif decorator notify")
	n.notifier.Notify()
}