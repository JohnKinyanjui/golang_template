package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
	"unicode"
)

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

// Generate unique name by appending a timestamp
func GenerateUniqueName(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixNano())
}

func Base64(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func ConvertSnakeCaseToTitleCase(value string) string {
	words := strings.Split(value, "_")

	for i, word := range words {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func FormatPrice(price float64) string {
	return fmt.Sprintf("%.0f", price)
}

func FormatInt(value int) string {
	return fmt.Sprintf("%d", value)
}
