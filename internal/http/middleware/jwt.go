package middleware

import (
	"HabitFlow/internal/domain/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey []byte

func GenerateToken(user_id int, admin bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	Claims := &models.Claims{
		User_id: user_id,
		Admin:   admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenStting, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return tokenStting, nil
}
func ValidateToken(tokenstr string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenstr, claims, func(token *jwt.Token) (interface{}, error) { return JwtKey, nil })
	if err != nil {
		return nil, err

	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil

}
