package interfaces

import (
	"fmt"
	"log"
)

type notifier interface {
	send(message, recipient string)
}

type emailSender struct {
	senderEmail string
}

func (en *emailSender) send(message, recipient string) {
	// logica pra enviar por email
	fmt.Printf("sent '%v'to %v via email.", message, recipient)
}

type SMSSender struct {
	senderNumber string
}

func (sms *SMSSender) send(message, recipient string) {
	// logica pra enviar por sms
	fmt.Printf("sent '%v'to %v via sms.", message, recipient)
}

type NotificationService struct {
	sender notifier
}

func (ns *NotificationService) SendNotification(message, recipient string) {
	log.Println("sending notification")
	ns.sender.send("enviando do service", "arthur")
}

func main() {

}
