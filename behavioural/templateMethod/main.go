package main

import "fmt"

func main() {
	sms := &SMSOTP{}
	otpService := &otp{interfaceOTP: sms}
	otpService.GenAndSendOTP(1)
	fmt.Println()
	email := &EmailOTP{}
	otpService = &otp{interfaceOTP: email}
	otpService.GenAndSendOTP(2)

}
