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

func (o *otp) GenAndSendOTP(otpLength int) error {
	otpGenerated := o.interfaceOTP.genRandomOTP(otpLength)
	o.interfaceOTP.saveOTPCache(otpGenerated)
	msg := o.interfaceOTP.getMessage(otpGenerated)
	err := o.interfaceOTP.sendNotification(msg)
	if err != nil {
		return err
	}
	o.interfaceOTP.publishMetric()
	return nil
}

