package util

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCmpText(t *testing.T) {
	text := "hello world!"
	encryptText, err := EncryptText(text)
	require.NoError(t, err)
	fmt.Println(encryptText)
	decryptText, err := DecryptText(encryptText)
	require.Equal(t, decryptText, text)
}
