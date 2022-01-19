package service

import (
	"log"
	"net/smtp"
)

func SendMailNetSMTP(mailto []string) {

	from := "khilmiaminudin715@gmail.com"
	password := "cxxapfaatxpyoalv"

	to := mailto
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("My super secret message.")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}

}
