package main

import "fmt"

type SMSOTP struct{
	otp
}

func (o *SMSOTP) genRandomOTP(n int) string {
	fmt.Println("SMS genRandom OTP")
	return ""
}

func (o *SMSOTP) saveOTPCache(val string) {
	fmt.Println("SMS saveOTPCache val:", val)
}

func (o *SMSOTP) getMessage(msg string) string {
	fmt.Println("SMS getMessage:", msg)
	return ""
}

func (o *SMSOTP) sendNotification(notif string) error {
	fmt.Println("SMS notif: ", notif)
	return nil
}

func (o *SMSOTP) publishMetric() {
	fmt.Println("SMS publish Metric")
}
