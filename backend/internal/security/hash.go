package security

import "golang.org/x/crypto/bcrypt"

func HashCode(code string) (string, error) {
	codeHash, err := bcrypt.GenerateFromPassword([]byte(code), 10)
	if err != nil {
		return "", err
	}
	return string(codeHash), nil
}
