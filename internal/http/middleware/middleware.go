package middleware

import (
	"HabitFlow/internal/auth"
	"net/http"
)

type AuthMiddleware struct {
	jwttoken auth.JWT
}

func NewAuthMiddleware(jwttoken auth.JWT) *AuthMiddleware {
	return &AuthMiddleware{jwttoken: jwttoken}
}

func (au *AuthMiddleware) AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		claims, err := au.jwttoken.ValidateToken(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if !claims.Admin {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next(w, r)
	}

}
