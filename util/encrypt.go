package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	AESKey = "32-byte-long-key-123456789012345"
)

func EncryptText(text string) (string, error) {
	key := []byte(AESKey)
	byteText := []byte(text)

	ciphertext, err := encryptAES(key, byteText)
	if err != nil {
		return "", fmt.Errorf("encrypt errror:%w", err)

	}
	return hex.EncodeToString(ciphertext), nil
}

func DecryptText(ciphertext string) (string, error) {
	key := []byte(AESKey)
	decryptText, err := decryptAES(key, []byte(ciphertext))
	if err != nil {
		return "", fmt.Errorf("decrypt errror:%w", err)
	}
	return string(decryptText), nil
}

func encryptAES(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func decryptAES(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("密文太短")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
