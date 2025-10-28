package notification

import "fmt"

type Notifier interface {
    Send(message string) string
}

type EmailNotifier struct{}
type SMSNotifier struct{}

func (e EmailNotifier) Send(message string) string {
    return fmt.Sprintf("[EMAIL SENT]: %s", message)
}

func (s SMSNotifier) Send(message string) string {
    return fmt.Sprintf("[SMS SENT]: %s", message)
}
