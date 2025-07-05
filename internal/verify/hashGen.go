package verify

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func generateHash() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal("Ошибка генерации случайных байт:", err)
	}
	return hex.EncodeToString(bytes)

}
