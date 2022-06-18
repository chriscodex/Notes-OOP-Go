package main

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
