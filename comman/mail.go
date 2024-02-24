package comman

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to string, body string, subject string) (bool, error) {
	emailStatus := false

	err := Getenv()
	if err != nil {
		log.Print("error loading .env file" + err.Error())
	}
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	auth := smtp.PlainAuth("", from, password, smtpServer)

	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	err = smtp.SendMail(fmt.Sprintf("%s:%s", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return false, err
	}

	fmt.Println("Email sent successfully!")

	return emailStatus, nil
}
