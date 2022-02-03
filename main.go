package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	fmt.Println("Golang email app running")
	email()
}

func email() {
	// sender data
	from := os.Getenv("user") //ex: "teste01@gmail.com"
	password := os.Getenv("pass")   // ex: "teste01@"
	// receiver address
	toEmail := os.Getenv("userResp") // ex: "resp01@gmail.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: Our Golang Email\n"
	body := "our first email!"
	message := []byte(subject + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// send mail
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}