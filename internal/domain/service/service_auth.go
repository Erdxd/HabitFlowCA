package service

import (
	"HabitFlow/internal/auth"
	"log"
)

type TokenService struct {
	jwt auth.JWT
}

func NewTokenService(jwt auth.JWT) *TokenService {
	return &TokenService{jwt: jwt}
}
func (TS *TokenService) CreateToken(user_id int, admin bool) (string, error) {
	token, err := TS.jwt.GenerateToken(user_id, admin)
	if err != nil {
		log.Println(err)
		return "", err

	}
	return token, nil

}
