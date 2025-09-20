package api

import "fmt"

func SendSMS(phone string, message string) {
	fmt.Println("Отправлено СМС на %s с сообщением %s", phone, message)
}
