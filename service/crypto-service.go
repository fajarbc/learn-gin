package service

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
// const SecretKey string = "my_secret_key"

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return data
}

func Encrypt(text, SecretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := MakeByte(plainText)
	cfb.XORKeyStream(cipherText, plainText)

	return Encode(cipherText), nil
}

func Decrypt(text, SecretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := MakeByte(cipherText)
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func MakeByte(b []byte) []byte {
	return make([]byte, len(b))
}
