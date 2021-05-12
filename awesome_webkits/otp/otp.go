package otp

import (
	otpgmail "awesome_webkits/otp/gmail"

	"fmt"
)

type EmailTokenModel struct {
	Email string
	Token string
}

func (e *EmailTokenModel) SendActivationToken() error {
	body := fmt.Sprintf("Your activation token is : %s", e.Token)

	return otpgmail.SendActivationToken(e.Email, "Authentication Token", body)
}

func (e *EmailTokenModel) SendResetPassToken() error {
	body := fmt.Sprintf("Your reset password token is : %s", e.Token)

	return otpgmail.SendActivationToken(e.Email, "Reset Password Token", body)
}
