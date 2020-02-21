package main

type Iotp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type otp struct{
	interfaceOTP Iotp
}

func (o *otp) genRandomOTP(n int) string {
	return o.interfaceOTP.genRandomOTP(n)
}

func (o *otp) saveOTPCache(val string) {
	o.interfaceOTP.saveOTPCache(val)
}

func (o *otp) getMessage(msg string) string {
	return o.interfaceOTP.getMessage(msg)
}

func (o *otp) sendNotification(notif string) error {
	return o.interfaceOTP.sendNotification(notif)
}

func (o *otp) publishMetric() {
	o.interfaceOTP.publishMetric()
}

