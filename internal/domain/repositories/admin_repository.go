package repositories

import "HabitFlow/internal/domain/models"

type Admin interface {
	GetDataAboutAllUsers() ([]models.User, error)
	ChangePasswordForUser(id_user int, newpasswordhashed string) error
	DeleteAccount(id_user int) error
	DeleteAllHabits(user_id int) error
}
