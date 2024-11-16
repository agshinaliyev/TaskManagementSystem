package implementing

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct {
}

type SMSNotifier struct {
}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("Sending Email Notification : ", message)
}

func (s SMSNotifier) Notify(message string) {
	fmt.Println("Sendin SMS notification : ", message)
}

func SendMessage(notifier Notifier, message string) {

	notifier.Notify(message)
}
