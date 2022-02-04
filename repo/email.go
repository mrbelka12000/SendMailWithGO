package repo

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/smtp"
)

//DeliverEmail отправляем сообщение по указанному адресу
func (c *Confirm) DeliverEmail(receiver string) error {
	code := generateConfirm()
	cfg := NewConfig(receiver)
	auth := smtp.PlainAuth("", cfg.From, cfg.Password, cfg.Host)

	msg := fmt.Sprintf(`
	Hi there, you are new on my site
	so you need to confrim your email
	please click on buttom link to go
	to acssess page
	http://localhost:8080/receive?confrim=%v `, code)

	err := smtp.SendMail(cfg.Addres, auth, cfg.From, []string{receiver}, []byte(msg))
	if err != nil {
		return err
	}
	c.Insert(receiver, code)
	return nil
}

func generateConfirm() string {
	return uuid.NewV4().String()
}
