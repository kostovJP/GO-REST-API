package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//[]byte() is used to convert something into a sequence of bytes. 
	// bcrypt.GenerateFromPassword() wants a sequence of bytes as first argument
	// so we have to first conver the string into a sequence of bytes.
	// later we can convert it back to string by using string(bytes)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	hashedPassword := string(bytes)
	return hashedPassword, err
}