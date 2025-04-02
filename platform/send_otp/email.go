package send_otp

import (
	"god/pkg/config"
	"gopkg.in/gomail.v2"
)

type SendOtpEmail struct {
	dialer *gomail.Dialer
	Email  string
	ApiKey string
}

func NewSendOtpEmail() *SendOtpEmail {
	email := config.StringEnv("SENDER_EMAIL")
	apiKey := config.StringEnv("API_KEY_EMAIL")
	host := config.StringEnv("SMTP_HOST")
	port := config.IntEnv("SMTP_PORT")
	return &SendOtpEmail{
		dialer: gomail.NewDialer(host, port, email, apiKey),
		Email:  email,
		ApiKey: apiKey,
	}
}

func (s *SendOtpEmail) SendOtp(toEmail, otp string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.Email)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", "Your OTP Code")
	mailer.SetBody("text/plain", "Your OTP is: "+otp)

	if err := s.dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}
