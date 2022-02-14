package email

import gomail "gopkg.in/mail.v2"

type Email struct {
	Emails   []string `json:"emails"`
	Subjects string   `json:"subjects"`
	Body     string   `json:"body"`
}

// new object
func NewEmail() *Email {
	return &Email{}
}

type MailSender struct {
	From   string
	Dailer *gomail.Dialer
}

func NewMailSender() *MailSender {
	return &MailSender{}
}

// methods
func (ms *MailSender) Send(emailChan chan Email) error {

	m := gomail.NewMessage()
	m.SetHeader("from", ms.From)
	for ec := range emailChan {
		m.SetHeader("Subject", ec.Subjects)
		m.SetBody("text/html", ec.Body)
		for _, to := range ec.Emails {
			m.SetHeader("To", to)
			if e := ms.Dailer.DialAndSend(m); e != nil {
				panic(e)
			}
		}
	}
	return nil
}
