package service_test

import (
	"github.com/stretchr/testify/require"
)

func (s *SuiteService) Test_LoginSuccess() {
	isAuthenticated, _, _ := s.authorService.Login(s.DB, s.authorValid.Username, s.authorValid.Password)

	require.True(s.T(), isAuthenticated, "test login-invalid gagal")
}

func (s *SuiteService) Test_LoginFailed() {
	isAuthenticated, _, _ := s.authorService.Login(s.DB, s.authorInvalid.Username, s.authorInvalid.Password)

	require.False(s.T(), isAuthenticated, "test login-invalid gagal")
}
