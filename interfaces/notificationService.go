package main

import "fmt"

type Notification interface {
	Send(msg string) error
}

type Email struct {
}

type SMS struct {
}

func (e Email) Send(msg string) error {
	fmt.Println("Email sent: ", msg)
	return nil
}

func (s SMS) Send(msg string) error {
	fmt.Println("SMS sent: ", msg)
	return nil
}

func NotifyAll(notifier []Notification, msg string) {
	for _, n := range notifier {
		n.Send(msg)
	}
}

func NotificationService() {
	e := Email{}
	s := SMS{}
	NotifyAll([]Notification{e, s}, "Hello")
}
