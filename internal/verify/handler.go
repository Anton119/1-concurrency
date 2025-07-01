package verify

import (
	"email/configs"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func NewHandler(router *http.ServeMux, conf *configs.Config) {
	router.HandleFunc("/verify/{hash}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "hash")
	})
	router.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		if conf.Password == "fake_pass" {
			fmt.Fprintln(w, "Заглушка: письмо отправлено")
			log.Println("Заглушка: письмо отправлено")
			return
		}

		host := "smtp.yandex.ru"
		e := email.NewEmail()
		e.From = fmt.Sprintf("Anton Ivanov <%s>", conf.Email)
		e.To = []string{"t4gan20011002@yandex.ru"}
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		err := e.Send(conf.Address, smtp.PlainAuth("", conf.Email, conf.Password, host))
		if err != nil {
			http.Error(w, "Не удалось отправить письмо: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Письмо успешно отправлено")

	})
}
