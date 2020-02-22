package main

import "fmt"

type EmailOTP struct{
	otp
}

func (o *EmailOTP) genRandomOTP(n int) string {
	fmt.Println("Email genRandom OTP")
	return ""
}

func (o *EmailOTP) saveOTPCache(val string) {
	fmt.Println("Email saveOTPCache val:", val)
}

func (o *EmailOTP) getMessage(msg string) string {
	fmt.Println("Email getMessage:", msg)
	return ""
}

func (o *EmailOTP) sendNotification(notif string) error {
	fmt.Println("Email notif: ", notif)
	return nil
}

func (o *EmailOTP) publishMetric() {
	fmt.Println("SMS publish Metric")
}
