package service_test

import (
	"testing"

	"github.com/fajarbc/learn-gin/service"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	result, err := service.Login("user", "user")
	if err != nil {
		t.Error("Error Login")
	}
	require.True(t, result, "")
}
