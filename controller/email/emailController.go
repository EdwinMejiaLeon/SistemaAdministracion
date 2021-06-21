package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
)

type Dest struct {
	Name string
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func SendEmail(nameUser string, email string) {
	// nombre correo electronico - destinatario
	from := mail.Address{"Edwin Mejia", "em97759@gmail.com"}
	to := mail.Address{nameUser, email}
	subject := "Notificación administrador"
	dest := Dest{Name: to.Address}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("./email/email.html")
	checkErr(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkErr(err)

	message += buf.String()

	servername := "smtp.gmail.com:465"
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", "em97759@gmail.com", "1111204157", host)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkErr(err)

	client, err := smtp.NewClient(conn, host)
	checkErr(err)

	err = client.Auth(auth)
	checkErr(err)

	err = client.Mail(from.Address)
	checkErr(err)

	err = client.Rcpt(to.Address)
	checkErr(err)

	w, err := client.Data()
	checkErr(err)

	_, err = w.Write([]byte(message))
	checkErr(err)

	err = w.Close()
	checkErr(err)

	client.Quit()
}
