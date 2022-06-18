package main

// Factory Notification
type IntNotificationFactory interface {
	SendNotification()
	GetSender() ISender
}
