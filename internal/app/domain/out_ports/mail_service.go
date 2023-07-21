package out_ports

type MailService interface {
	SendMessage(to string, from string, subject string, text string) error
}
