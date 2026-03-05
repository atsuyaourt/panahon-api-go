package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyTokenOK(t *testing.T) {
	tok := NewAPIToken("")

	err := tok.VerifyToken()
	require.NoError(t, err)
}

func TestVerifyTokenWrongPrefix(t *testing.T) {
	tok := NewAPIToken(RandomString(32))

	err := tok.VerifyToken()
	require.Error(t, err)
}
