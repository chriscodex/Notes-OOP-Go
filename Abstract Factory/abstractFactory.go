package main

import "fmt"

// Factory Notification
type IntNotificationFactory interface {
	SendNotification()
	GetSender() IntSenderFactory
}

// Factory Sender
type IntSenderFactory interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// SMS Notifications
type SMSNotification struct {
}

// Methods of SMSNotification
func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification SMS")
}
func (SMSNotification) GetSender() IntSenderFactory {
	return SMSNotificationSender{}
}

// SMS Notification Sender
type SMSNotificationSender struct {
}

// Methods of SMSNotificationSender
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Email Notification
type EmailNotification struct {
}

// Methods of Email Notification
func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification Email")
}
func (EmailNotification) GetSender() IntSenderFactory {
	return EmailNotificationSender{}
}

// Email Notification Sender
type EmailNotificationSender struct {
}

// Methods of EmailNotificationSender
func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "Gmail"
}

//
func getNotificationFactory(notificationType string) (IntNotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Type")
}

//
func sendNotification(f IntNotificationFactory) {
	f.SendNotification()
}

func getMethod(f IntNotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	sms, _ := getNotificationFactory("SMS")
	email, _ := getNotificationFactory("Email")
	sendNotification(sms)
	sendNotification(email)

	getMethod(sms)
	getMethod(email)
}
