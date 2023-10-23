package middleware

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func SendEmailPenngajuan(img string, to string, url string, name string) error {
	var body bytes.Buffer
	t, err := template.ParseFiles("C:/Users/LENOVO/go/src/findnseek/middleware/template.html")
	if err != nil {
		logrus.Error("Failed to parse template:", err)
		return err
	}

	data := struct {
		Name string
		URL  string
		Img  string
	}{
		Name: name,
		URL:  url,
		Img:  img,
	}

	if err := t.Execute(&body, data); err != nil {
		logrus.Error("Failed to execute template:", err)
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "gempar835@gmail.com")
	m.SetHeader("To", to)
	m.SetAddressHeader("Cc", "gempar835@gmail.com", "Dan")
	m.SetHeader("Subject", "Find N Seek!")
	m.SetBody("text/html", body.String())
	// m.Attach("C:/Users/LENOVO/go/src/findnseek/middleware/Naruto.webp")

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
	m.SetBody("text/html", fmt.Sprintf("<h2>Hei Ternyata Benar Yang Kau Temui %s silahkan hubungi nomor nya ya %s</h2>", name, wa))

	d := gomail.NewDialer("smtp.gmail.com", 587, "gempar835@gmail.com", "xqrfdvnmscimxsgs")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return nil

}
