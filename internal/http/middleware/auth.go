package middleware

import (
	"log"
	"net/http"
)

func Cookie_userID(w http.ResponseWriter, r *http.Request) (int, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		log.Println("You are havent cookie")
		return 0, err
	}

	claims, err := ValidateToken(cookie.Value)
	if err != nil {
		log.Println("Your token is uncorrect")
		return 0, err
	}

	return claims.User_id, nil

}
