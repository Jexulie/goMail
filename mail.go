package mailer

import (
	"gopkg.in/gomail.v2"
)

// MailServer Struct
type MailServer struct {
	SMTP       string
	Port       int
	Email      string
	Password   string
	Attachment []string
}

// New SMPT
func New(SMTP string, Port int, Email, Password string) *MailServer {
	Attachment := []string{}
	return &MailServer{
		SMTP,
		Port,
		Email,
		Password,
		Attachment,
	}
}

// AddAttachment adds attachment
func (m *MailServer) AddAttachment(path string) {
	m.Attachment = append(m.Attachment, path)
}

// SendMail sends mail
func (m *MailServer) SendMail(recipients []string, subject, message string) error {

	z := gomail.NewMessage()
	z.SetHeader("From", m.Email)
	z.SetHeader("To", recipients...)
	z.SetHeader("Subject", subject)

	if len(m.Attachment) > 0 {
		for _, a := range m.Attachment {
			z.Attach(a)
		}
	}

	z.SetBody("text/html", message)

	s := gomail.NewDialer(
		m.SMTP,
		m.Port,
		m.Email,
		m.Password,
	)

	if err := s.DialAndSend(z); err != nil {
		return err
	}

	return nil
}
