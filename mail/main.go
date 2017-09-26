package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "mert.acel@ardictech.com", "Ma123456", "mail.ardictech.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"mert.acel@ardictech.com"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("mail.ardictech.com:443", auth, "mert.acel@ardictech.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
