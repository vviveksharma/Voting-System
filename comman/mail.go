package comman

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendEmail(to string, body string, subject string) (bool, error) {
	emailStatus := false

	err := godotenv.Load()
	if err != nil {
		log.Print("error loading .env file")
	}
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpServer := "smtp.gmail.com"
	smtpPort := 587
	auth := smtp.PlainAuth("", from, password, smtpServer)

	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		return false, err
	}

	fmt.Println("Email sent successfully!")

	return emailStatus, nil
}
