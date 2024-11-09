package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// EncryptMap encrypts a map[string]interface{} using XOR with a key and returns a base64 encoded string
func EncryptMap(data map[string]interface{}, key string) (string, error) {
	// Convert map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Encrypt the JSON data
	encrypted := make([]byte, len(jsonData))
	for i := 0; i < len(jsonData); i++ {
		encrypted[i] = jsonData[i] ^ key[i%len(key)]
	}

	// Encode the encrypted data to base64
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// DecryptMap decrypts a base64 encoded string back to a map[string]interface{}
func DecryptMap(encryptedStr string, key string) (map[string]interface{}, error) {
	// Decode the base64 string
	encrypted, err := base64.StdEncoding.DecodeString(encryptedStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %v", err)
	}

	// Decrypt the data
	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i++ {
		decrypted[i] = encrypted[i] ^ key[i%len(key)]
	}

	// Convert JSON back to map
	var result map[string]interface{}
	err = json.Unmarshal(decrypted, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return result, nil
}
