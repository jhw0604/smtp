package office365

import (
	"errors"
	"net/smtp"
)

func StartTlsAuth(username, password string) smtp.Auth {
	return &startTlsAuth{username: username, password: password}

}

type startTlsAuth struct {
	username string
	password string
}

func (s *startTlsAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(s.username), nil
}
func (s *startTlsAuth) Next(fromServer []byte, more bool) ([]byte, error) {
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
