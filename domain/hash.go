package domain

import "golang.org/x/crypto/bcrypt"

func Hash(b []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(b, cost)
}

func VerifyHash(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
