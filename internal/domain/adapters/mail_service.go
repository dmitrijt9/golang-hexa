package adapters

type MailService interface {
	SendMessage(to string, from string, subject string, text string) error
}
