package util

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCmpText(t *testing.T) {
	text := "hello world!"
	encryptText, err := EncryptAES(text)
	require.NoError(t, err)
	fmt.Println(encryptText)
	decryptText, err := DecryptAES(encryptText)
	require.Equal(t, decryptText, text)
}
