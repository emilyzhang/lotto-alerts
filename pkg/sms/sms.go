package sms

import(
	textmagic "github.com/textmagic/textmagic-rest-go"
)

func send() {
	client, err := sms.createClient("smtp.gmail.com", 587, "tester@gmail.com", "test")
}
