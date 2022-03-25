package service_test

import (
	"fmt"
	"testing"

	"github.com/fajarbc/learn-gin/service"
	"github.com/stretchr/testify/require"
)

const SecretKey string = "my_secret_key45678901234" // 24 chars

var plainText string = "Test Passed"
var base64Text string = "VGVzdCBQYXNzZWQ="
var encryptedText string = "iga/KLS2pGoiQWA="

func TestMain(m *testing.M) {
	fmt.Println("service test started")
	m.Run()
	fmt.Println("service test finished")
}

func TestEncode(t *testing.T) {
	result := service.Encode([]byte(plainText))
	require.Equal(t, base64Text, result, "test encode gagal")
}

func TestDecode(t *testing.T) {
	result := string(service.Decode(base64Text))
	require.Equal(t, plainText, result, "test decode gagal")
}

func TestEncrypt(t *testing.T) {
	result, err := service.Encrypt(plainText, SecretKey)
	if err != nil {
		t.Error("Error Encrypt")
	}
	require.Equal(t, encryptedText, result, "test encrypt gagal")
}

func TestDecrypt(t *testing.T) {
	result, err := service.Decrypt(encryptedText, SecretKey)
	if err != nil {
		t.Error("Error Decrypt")
	}
	require.Equal(t, plainText, result, "test decrypt gagal")
}
