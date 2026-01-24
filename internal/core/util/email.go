package util

import (
	"fmt"
	"log"

	"github.com/yehezkiel1086/go-gin-nextjs-auth/internal/adapter/config"
	"gopkg.in/gomail.v2"
)

func SendConfirmationEmail(httpConf *config.HTTP, email, token string, duration string) error {
	// generate confirmation url
	url := fmt.Sprintf("http://%s:%s/api/v1/confirm?token=%s", httpConf.Host, httpConf.Port, token)

	log.Printf("[+] Sending confirmation link %s to %s", url, email)

	// add gomail to send confirmation link to email
	m := gomail.NewMessage()

	m.SetHeader("From", httpConf.Email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "User registration confirmation")

	m.SetBody("text/html", `
		<h2>Email Confirmation</h2>
		<p>Please confirm your email by clicking the link below:</p>
		<p>
			<a href="`+url+`">Confirm Email</a>
		</p>
		<p>This link will expire in ` + duration + ` minutes.</p>
	`)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		httpConf.Email,
		httpConf.EmailPassword,
	)

	d.SSL = false // STARTTLS

	return d.DialAndSend(m)
}
