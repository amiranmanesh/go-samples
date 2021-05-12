package otp_gmail

import (
	"awesome_webkits/utils/env"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
	"strconv"
)

var (

	// Sender data.
	//todo
	from = env.GetEnvItem("EMAIL_ADDRESS")
	//password = env.GetEnvItem("EMAIL_PASSWORD")
	password = "pass"

	// smtp server configuration.
	smtpHost    = env.GetEnvItem("EMAIL_HOST")
	smtpPort, _ = strconv.Atoi(env.GetEnvItem("EMAIL_PORT"))
)

func SendActivationToken(to, subject, body string) error {
	m := gomail.NewMessage()
	// Set E-Mail sender
	m.SetHeader("From", from)
	// Set E-Mail receivers
	m.SetHeader("To", to)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	//body := fmt.Sprintf("Your activation token is : %s", a.token)
	m.SetBody("text/plain", body)

	// Settings for SMTP server
	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	return d.DialAndSend(m)
}

/*
func sendHTML() {

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, _ := template.ParseFiles("temp.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	_ = t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+string(smtpPort), auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
*/
