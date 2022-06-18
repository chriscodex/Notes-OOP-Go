package main

import "fmt"

// Factory Notification
type IntNotificationFactory interface {
	SendNotification()
	GetSender() IntSender
}

// Factory Sender
type IntSender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

// Notifications
type SMSNotification struct {
}

// Methods of SMSNotification
func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification SMS")
}

// SMSNotificationSender
type SMSNotificationSender struct {
}

// Methods of SMSNotificationSender
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
