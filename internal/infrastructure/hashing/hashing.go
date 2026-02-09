package hashing

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type HashService struct {
}

func NewHashService() *HashService {
	return &HashService{}
}
func (HS *HashService) Hash(password string) (string, error) {
	HasgedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to encrypt your password")
		return "", err
	}
	return string(HasgedPassword), nil

}
func (HS *HashService) Compare(HashedPassword, Password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPassword), []byte(Password))
	return err == nil

}
