package auth

import "HabitFlow/internal/domain/models"

type JWT interface {
	GenerateToken(user_id int, admin bool) (string, error)
	ValidateToken(tokenstr string) (*models.Claims, error)
}
