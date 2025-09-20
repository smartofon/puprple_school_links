package secure

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSecureID(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return hex.EncodeToString(b), err
}

func GenerateSecureDigits(n int) (string, error) {
	const charset = "0123456789"
	b := make([]byte, n)
	_, err := rand.Read(b) // Заполняем байты случайными данными
	if err != nil {
		return "", err
	}

	for i := range b {
		b[i] = charset[b[i]%10] // Ограничиваем диапазон до 10
	}

	return string(b), nil
}
