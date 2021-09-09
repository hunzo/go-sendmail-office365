package main

import (
	"fmt"
	"go-smtp/service"
	"log"
)

func main() {
	m := service.Mail{
		FromEmailAdddress: "EMAIL_ADDRESS",
		FromName:          "EMAIL_NAME",
		FromPassword:      "EMAIL_PASSWORD",
		ToEmailAddress:    "RECIPIENT_ADDRESS",
	}

	m.Subject = "This Subject"
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\nFrom: " + m.FromName + " <" + m.FromEmailAdddress + ">" + "\n"

	m.Body.Write([]byte(
		fmt.Sprintf("Subject: %s\n%s\n\n This is Body", m.Subject, mimeHeaders)))

	err := service.Mailer(m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Mail sent to: %s, subject: %s\n", m.FromEmailAdddress, m.Subject)

}
