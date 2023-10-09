package services

import (
	"hexa-example-go/internal/app/domain/out_ports"
	"hexa-example-go/internal/logger"
	"net/smtp"
)

type MailService struct {
	host string
	port string
}

func NewMailService(logger logger.Logger, host string, port string) out_ports.MailService {
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
