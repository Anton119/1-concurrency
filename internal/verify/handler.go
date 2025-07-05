package verify

import (
	"email/configs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

const verificationFile = "pending.json"

type SendRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewHandler(router *http.ServeMux, conf *configs.Config) {
	router.HandleFunc("/verify/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		hash := strings.TrimPrefix(r.URL.Path, "/verify/")
		if hash == "" {
			http.Error(w, "Хеш не указан", http.StatusBadRequest)
			return
		}

		data, err := loadVerificationData(verificationFile)
		if err != nil {
			http.Error(w, "Ошибка загрузки данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		emailAddr, ok := data[hash]
		if !ok {
			http.Error(w, "Неверная или просроченная ссылка", http.StatusBadRequest)
			return
		}

		delete(data, hash)
		err = saveVerificationData(verificationFile, data)
		if err != nil {
			http.Error(w, "Ошибка сохранения данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Email %s успешно подтверждён!", emailAddr)
	})

	router.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var req SendRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Неверный формат JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		if req.Password == "fake_pass" {
			fmt.Fprintln(w, "Заглушка: письмо отправлено")
			log.Println("Заглушка: письмо отправлено")
			return
		}

		// Загружаем текущие данные
		data, err := loadVerificationData(verificationFile)
		if err != nil {
			http.Error(w, "Ошибка загрузки данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Генерируем хеш и сохраняем связку
		hash := generateHash()
		data[hash] = req.Email

		err = saveVerificationData(verificationFile, data)
		if err != nil {
			http.Error(w, "Ошибка сохранения данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		e := email.NewEmail()
		e.From = fmt.Sprintf("Anton Ivanov <%s>", req.Email)
		e.To = []string{"t4gan20011002@yandex.ru"} // TODO: заменить на динамического получателя
		e.Subject = "Подтверждение регистрации"
		e.Text = []byte(fmt.Sprintf("Перейдите по ссылке для подтверждения: http://localhost:8081/verify/%s", hash))

		host := "smtp.yandex.ru"
		err = e.Send(host+":587", smtp.PlainAuth("", req.Email, req.Password, host))
		if err != nil {
			http.Error(w, "Не удалось отправить письмо: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "всё дошло до отправки")

	})
}

// получили email
// сгенерировали hash
// сохранили email с хешом
// отправили письмо содержащее ссылку с хэшом
