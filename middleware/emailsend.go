package middleware

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmailPenngajuan(img string, to string, url string, name string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "gempar835@gmail.com")
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", "gempar835@gmail.com", "Dan")
	m.SetHeader("Subject", "Find N Seek!")
	m.SetBody("text/html", fmt.Sprintf("<h2>Hei Apakah Ini Barang Miiikmu?</h2> <img src=%s>", img))
	// m.Attach(img)

	d := gomail.NewDialer("smtp.gmail.com", 587, "gempar835@gmail.com", "xqrfdvnmscimxsgs")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil

}
func SendEmailPersetujuan(to string, wa string, name string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "gempar835@gmail.com")
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", "gempar835@gmail.com", "Dan")
	m.SetHeader("Subject", "Find N Seek!")
	m.SetBody("text/html", fmt.Sprintf("<h2>Hei Ternyata Benar Yang Kau Temui milik, %s silahkan hubungi nomor nya ya : %s</h2>", name, wa))

	d := gomail.NewDialer("smtp.gmail.com", 587, "gempar835@gmail.com", "xqrfdvnmscimxsgs")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil

}
