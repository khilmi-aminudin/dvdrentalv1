package service

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
	"gopkg.in/gomail.v2"
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

func SendMailGoMail() {

	htmlbody := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Hello World</title>
	</head>
	<body>
		<h1>Test Email</h1>
		<h5>Inini adalah email untuk test</h5>
	</body>
	</html>`

	to := "rentalboygame254@gmail.com"
	cc := "khilmi_aminudin@yahoo.co.id"

	err := godotenv.Load()
	helper.PanicIfError(err)

	sender := os.Getenv("MAIL_USERNAME")

	var (
		smtpHost         = os.Getenv("MAIL_HOST")
		smtpPort         = os.Getenv("MAIL_PORT")
		smtpAuthUsername = os.Getenv("MAIL_USERNAME")
		smtpAuthPassword = os.Getenv("MAIL_PASSWORD")
	)

	smtpport, _ := strconv.Atoi(smtpPort)

	configSender := fmt.Sprintf("PT. Sumber Jaya <%s>", sender)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", configSender)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Cc", cc)
	mailer.SetHeader("Subject", "Test Mail")
	mailer.SetBody("text/html", htmlbody)
	mailer.Attach("example-logo.jpg")

	dialer := gomail.NewDialer(
		smtpHost,
		smtpport,
		smtpAuthUsername,
		smtpAuthPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Mail Sent")

}
