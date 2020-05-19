package office365

import (
	"errors"
	"net/smtp"
)

//StartTLSAuth is STARTTLS
func StartTLSAuth(username, password string) smtp.Auth {
	return &startTlsAuth{username: username, password: password}

}

type startTLSAuth struct {
	username string
	password string
}

func (s *startTLSAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(s.username), nil
}
func (s *startTLSAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(s.username), nil
		case "Password:":
			return []byte(s.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}
