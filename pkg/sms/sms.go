package sms

import (
	"crypto/tls"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

// Send sends messages to the number specified (must be an AT&T number).
func Send(body string, num string) {
	d := gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("GMAIL_USER"), os.Getenv("PASSWORD"))
	// Disabling TLS is insecure - should be not used in production!
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	s, err := d.Dial()
	if err != nil {
		fmt.Println("Could not create NewDialer", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "alerts@lotto.com")
	m.SetAddressHeader("To", num+"@txt.att.net", "")
	m.SetBody("text/html", body)

	if err := gomail.Send(s, m); err != nil {
		fmt.Println("Could not send email", err)
	}
	m.Reset()
}
