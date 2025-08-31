package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(s string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	return string(byte), err
}

func CheckPasswordHash(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
