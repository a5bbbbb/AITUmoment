package service

type EmailTransmitter interface {
	Send(emailAddress, subject, body string) error
}
