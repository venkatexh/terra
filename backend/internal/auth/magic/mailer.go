package magic

import "log"

type Mailer interface {
	SendMagicLink(email, link string) error
}

type ConsoleMailer struct{}

func (c ConsoleMailer) SendMagicLink(email, link string) error {
	log.Printf("Sending magic link to %s: %s", email, link)
	return nil
}