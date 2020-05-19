package office365

import "strings"

//Msg is ..
type Msg interface {
	From() string
	To() []string
	Byte() []byte
}

type message struct {
	subject     string
	from        string
	to          []string
	cc          []string
	bcc         []string
	contentType string
	body        string
}

//NewMessage is create message struct
func NewMessage(subject string, from string, to []string, cc []string, bcc []string, contentType string, body string) Msg {
	return message{subject: subject, from: from, to: to, cc: cc, bcc: bcc, contentType: contentType, body: body}
}

//From is return from email
func (m message) From() string {
	return m.from
}

//To is return list of recever emails
func (m message) To() []string {
	return append(append(m.to, m.cc...), m.bcc...)
}

//Byte is return convert messge to bytes
func (m message) Byte() []byte {
	var msg string
	if m.subject != "" {
		msg = msg + "Subject: " + m.subject + "\r\n"
	}
	if m.from != "" {
		msg = msg + "From: " + m.from + "\r\n"
	}
	if m.to != nil && len(m.to) > 0 {
		msg = msg + "To: " + strings.Join(m.to, ";") + "\r\n"
	}
	if m.cc != nil && len(m.cc) > 0 {
		msg = msg + "CC: " + strings.Join(m.cc, ";") + "\r\n"
	}
	if m.bcc != nil && len(m.bcc) > 0 {
		msg = msg + "Bcc: " + strings.Join(m.bcc, ";") + "\r\n"
	}
	if m.contentType != "" {
		msg = msg + "Content-Type: " + m.contentType + "\r\n"
	}
	if m.body != "" {
		msg = msg + "\r\n" + m.body + "\r\n"
	}

	return []byte(msg)
}
