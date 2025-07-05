package verify

import (
	"encoding/json"
	"os"
)

type VerificationData struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func saveVerificationData(filename string, data map[string]string) error {
	// 1. Открываем (создаём) файл для записи
	file, err := os.Create(filename)
	if err != nil {
		return err // если ошибка — возвращаем её вызывающему коду
	}
	// Обязательно закрываем файл после завершения функции
	defer file.Close()

	// 2. Создаём JSON-энкодер, который будет писать данные в файл
	encoder := json.NewEncoder(file)

	// 3. Задаём красивое форматирование JSON (отступы и переносы строк)
	encoder.SetIndent("", "  ")

	// 4. Кодируем мапу data в JSON и записываем в файл
	// Если кодирование/запись прошла успешно, возвращается nil
	// Иначе — ошибка
	return encoder.Encode(data)
}

func loadVerificationData(filename string) (map[string]string, error) {
	// 1. Пытаемся открыть файл с именем filename для чтения
	file, err := os.Open(filename)
	if err != nil {
		// 2. Если файл не найден (ошибка типа "файл не существует")
		if os.IsNotExist(err) {
			// Возвращаем новую пустую мапу и nil (нет ошибки)
			return make(map[string]string), nil
		}
		// Если другая ошибка — возвращаем nil и ошибку вызывающему коду
		return nil, err
	}
	// Обязательно закрываем файл после завершения функции
	defer file.Close()

	// 3. Объявляем переменную data для хранения результата (мапа)
	var data map[string]string

	// 4. Создаём JSON-декодер, который будет читать из файла
	decoder := json.NewDecoder(file)

	// 5. Декодируем JSON из файла в переменную data
	err = decoder.Decode(&data)
	if err != nil {
		// Если при декодировании произошла ошибка — возвращаем её
		return nil, err
	}

	// 6. Если всё успешно — возвращаем загруженную мапу и nil ошибки
	return data, nil
}
