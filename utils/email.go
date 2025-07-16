package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(receiverEmail, token string) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := os.Getenv("EMAIL_SENDER")
	senderPassword := os.Getenv("EMAIL_PASSWORD")
	host := "http://146.190.102.54:9702/"

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	subject := "Subject: Password Reset Request\r\n"
	headers := "From: Kukky App <" + senderEmail + ">\r\n" +
		"To: " + receiverEmail + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n"

	body := fmt.Sprintf(`We received a request to reset your password.

Use the following token to proceed with your password reset (Valid for 10 minutes):
%s/reset-password/%s

Kukky App Team
`, host, token)

	message := []byte(subject + headers + "\r\n" + body)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		senderEmail,
		[]string{receiverEmail},
		message,
	)
	if err != nil {
		return err
	}

	return nil
}
