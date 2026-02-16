package middleware

import (
	"HabitFlow/internal/auth"
	"net/http"
)

type JWTToken struct {
	jwttoken auth.JWT
}

func NewJwtKey(jwttoken auth.JWT) *JWTToken {
	return &JWTToken{jwttoken: jwttoken}
}

func (jw *JWTToken) Cookie_userID(w http.ResponseWriter, r *http.Request) (int, error) {
	cookie, err := r.Cookie("token")
	if err != nil {

		return 0, err
	}

	claims, err := jw.jwttoken.ValidateToken(cookie.Value)
	if err != nil {

		return 0, err
	}

	return claims.User_id, nil

}
