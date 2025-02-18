package tests

import (
	"os"
	"testing"

	"github.com/go-gomail/gomail"
	"github.com/stretchr/testify/assert"
)

type MockDialer struct{}

func (d *MockDialer) DialAndSend(m *gomail.Message) error {
	return nil
}

func TestSendEmail(t *testing.T) {
	os.Setenv("SMTP_USER", "test@example.com")
	os.Setenv("SMTP_HOST", "smtp.example.com")
	os.Setenv("SMTP_PASSWORD", "password")

	tests := []struct {
		name      string
		to        interface{}
		subject   string
		body      string
		expectErr bool
	}{
		{"ValidEmail", "recipient@example.com", "Test Subject", "Test Body", false},
		{"InvalidEmailType", 12345, "Test Subject", "Test Body", true},
		{"EmptyEmail", "", "Test Subject", "Test Body", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gomailDialer := &MockDialer{}
			m := gomail.NewMessage()
			m.SetHeader("From", os.Getenv("SMTP_USER"))
			m.SetHeader("To", tt.to.(string))
			m.SetHeader("Subject", tt.subject)
			m.SetBody("text/html", tt.body)

			err := gomailDialer.DialAndSend(m)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
