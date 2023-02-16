package cryptox

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(password string, salt string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
