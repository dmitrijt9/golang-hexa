package services

import (
	"net/smtp"

	"go.uber.org/zap"
)

type MailService struct {
	host string
	port string
}

func NewMailService(logger zap.Logger, host string, port string) *MailService {
	return &MailService{
		host: host,
		port: port,
	}
}

func (m *MailService) SendMessage(to string, from string, subject string, text string) error {

	toAddrs := []string{to}
	address := m.host + ":" + m.port
	message := []byte(subject + text)

	err := smtp.SendMail(address, nil, from, toAddrs, message)
	if err != nil {
		return err
	}

	return nil
}
