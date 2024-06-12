package main

import (
	"fmt"
	"net/smtp"

	"github.com/sfreiberg/gotwilio"
)

// Notifier interface
type Notifier interface {
	Notify(message string) error
}

// EmailNotifier struct
type EmailNotifier struct {
	EmailAddress string
}

// Implementasi metode Notify pada EmailNotifier
func (en EmailNotifier) Notify(message string) error {
	// Setup SMTP server configuration.
	// Ganti konfigurasi sesuai server SMTP Anda
	smtpHost := "smtp.example.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", "your-email@example.com", "your-email-password", smtpHost)

	to := []string{en.EmailAddress}
	msg := []byte("To: " + en.EmailAddress + "\r\n" +
		"Subject: Notification\r\n" +
		"\r\n" +
		message + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, "your-email@example.com", to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Printf("Email sent to %s: %s\n", en.EmailAddress, message)
	return nil
}

// SMSNotifier struct
type SMSNotifier struct {
	PhoneNumber string
}

// Implementasi metode Notify pada SMSNotifier
func (sn SMSNotifier) Notify(message string) error {
	accountSid := "your_account_sid"
	authToken := "your_auth_token"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := "+1234567890" // Ganti dengan nomor Twilio Anda
	to := sn.PhoneNumber

	_, _, err := twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		return fmt.Errorf("failed to send SMS: %v", err)
	}

	fmt.Printf("SMS sent to %s: %s\n", sn.PhoneNumber, message)
	return nil
}

// Fungsi untuk mengirim notifikasi menggunakan Notifier
func sendNotification(n Notifier, message string) {
	err := n.Notify(message)
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}

func main() {
	emailNotifier := EmailNotifier{EmailAddress: "abdulrosyidumar@gmail.com"}
	smsNotifier := SMSNotifier{PhoneNumber: "+6285939323270"}

	message := "Hello, this is a notification!"

	// Mengirim notifikasi menggunakan EmailNotifier
	sendNotification(emailNotifier, message)

	// Mengirim notifikasi menggunakan SMSNotifier
	sendNotification(smsNotifier, message)
}
