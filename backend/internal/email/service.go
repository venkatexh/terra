package email

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Service struct {
	apiKey string
	from   string
}

func NewService() *Service {
	return &Service{
		apiKey: os.Getenv("SENDGRID_API_KEY"),
		from:   os.Getenv("FROM_EMAIL"),
	}
}

func (s *Service) SendMagicLink(toEmail, link string) error {
	from := mail.NewEmail("Terra", s.from)
	to := mail.NewEmail("", toEmail)

	subject := "Your Terra login link"

	html := fmt.Sprintf(`
		<h2>Login to Terra</h2>
		<p>Click the link below to login:</p>
		<a href="%s">%s</a>
	`, link, link)

	message := mail.NewSingleEmail(from, subject, to, "", html)

	client := sendgrid.NewSendClient(s.apiKey)

	_, err := client.Send(message)

	if err != nil {
		return err
	}
	return nil
}
