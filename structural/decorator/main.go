package main

func main() {
	basicnotifier := &BasicNotifier{}
	enableFacebookNotif := true
	enableSMSNotif := true
	if enableFacebookNotif {
		basicnotifier = NewFacebookNotifDecorator(basicnotifier)
	}
}
