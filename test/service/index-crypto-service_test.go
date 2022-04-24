package service_test

import (
	"github.com/fajarbc/learn-gin/service"

	"github.com/stretchr/testify/require"
)

func (s *SuiteService) Test_Encode() {
	result := service.Encode([]byte(s.plainText))
	require.Equal(s.T(), s.base64Text, result, "test encode gagal")
}

func (s *SuiteService) Test_Decode() {
	result := string(service.Decode(s.base64Text))
	require.Equal(s.T(), s.plainText, result, "test decode gagal")
}

func (s *SuiteService) Test_Encrypt() {
	result, err := service.Encrypt(s.plainText, s.secretKey)
	if err != nil {
		s.T().Error("Error Encrypt")
	}
	require.Equal(s.T(), s.encryptedText, result, "test encrypt gagal")
}

func (s *SuiteService) Test_Decrypt() {
	result, err := service.Decrypt(s.encryptedText, s.secretKey)
	if err != nil {
		s.T().Error("Error Decrypt")
	}
	require.Equal(s.T(), s.plainText, result, "test decrypt gagal")
}
