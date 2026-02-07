package repositories

import "HabitFlow/internal/domain/models"

type Habit interface {
	CheckHabit(id_user int) ([]models.HabitFlow, error)
	AddHabit(Habits models.HabitFlow, Id_user int) error
	DeleteHabit(id_user, id int) error
	ChangeStatusHabit(user_id, streak, id int) error
	GetStreakHabit(user_id, id int) (int, error)
	ResetAllStatusHabit() error
	DeleteAllHabits(user_id int) error
}
