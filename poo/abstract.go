package main

import (
	"fmt"
	"strings"
)

type Sender interface {
	GetChannel() string
	GetMethod() string
}

type SMSNotificationSender struct {
	channel string
	method  string
}

func (s *SMSNotificationSender) GetChannel() string {
	return s.channel
}

func (s *SMSNotificationSender) GetMethod() string {
	return s.method
}

type EmailNotificationSender struct {
	channel string
	method  string
}

func (e *EmailNotificationSender) GetChannel() string {
	return e.channel
}

func (e *EmailNotificationSender) GetMethod() string {
	return e.method
}

type Notifier interface {
	Notify()
	GetSender() Sender
}

type SMSNotifier struct {
}

func (SMSNotifier) Notify() {
	println("Sending an SMS")
}

func (SMSNotifier) GetSender() Sender {
	return &SMSNotificationSender{channel: "SMS", method: "Twilio"}
}

type EmailNotifier struct {
}

func (EmailNotifier) Notify() {
	println("Sending an email")
}

func (EmailNotifier) GetSender() Sender {
	return &EmailNotificationSender{channel: "Email", method: "SMTP"}
}

func getNotifier(notifierType string) (Notifier, error) {
	_type := strings.ToLower(notifierType)

	if _type == "email" {
		return &EmailNotifier{}, nil
	} else if _type == "sms" {
		return &SMSNotifier{}, nil
	}

	return nil, fmt.Errorf("notifier type not found")
}

func main() {
	var notifier Notifier

	notifier, err := getNotifier("sms")

	if err == nil {
		notifier.Notify()
		println(notifier.GetSender().GetChannel())
		println(notifier.GetSender().GetMethod())
	} else {
		println("notifier for sms not found")
	}

	notifier, err = getNotifier("email")

	if err == nil {
		notifier.Notify()
		println(notifier.GetSender().GetChannel())
		println(notifier.GetSender().GetMethod())
	} else {
		println("notifier for email not found")
	}
}
